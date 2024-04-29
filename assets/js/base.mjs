class Config {
    constructor(mode, connection_mode) {
        this.mode = mode
        if(typeof connection_mode !== "undefined") {
            this.connection_mode = connection_mode
        } else {
            this.connection_mode = "socket"
        }

    }

    isSetting() {
        return this.mode === "setting";
    }
}

class Socket {
    constructor() {
        this.callback = function(message) { alert(message) }
        this.server = new WebSocket(window.socket_host);
        this.server.onopen = (e) => {
            // alert("[open] Connection established");
            // alert("Sending to server");
            this.server.send("ping");
        };

        this.server.onmessage = (event) => {
            console.log(`[message] Data received from server: ${event.data}`);
            this.callback(event.data);
        };

        this.server.onclose = (event) => {
            if (event.wasClean) {
                console.log(`[close] Connection closed cleanly, code=${event.code} reason=${event.reason}`);
            } else {
                // e.g. server process killed or network down
                // event.code is usually 1006 in this case
                console.log('[close] Connection died');
            }
        };

        this.server.onerror = (error) => {
            console.error(`[error] ${error.message}`);
        };
    }

    onmessage(callback) {
        this.callback = callback
    }

    emit(event, msg) {

    }
}

class Buttons {
    constructor() {
        this.buttons = [
            $('.addPoint[data-player="1"]'),
            $('.addPoint[data-player="2"]'),
            $('.decPoint[data-player="1"]'),
            $('.decPoint[data-player="2"]'),
            $('.ball[data-player="1"]'),
            $('.ball[data-player="2"]'),
            $('.save[data-player="1"]'),
            $('.save[data-player="2"]'),
            $('.reset'),
            $('.switch-mode')
        ];
        this.buttons.forEach(b => this.observe(b));
    }

    observe(p) {
        p.on('click', (event) => {
            event.preventDefault();
            let obj = $(event.target);
            if(obj.hasClass("reset")) {
                fetch('/point/reset').then(r => console.log(r));
            } else if(obj.hasClass("save")) {
                let playerNo = obj.data('player')
                let newName = $('#player' + playerNo).find('.name').val();
                const data = new URLSearchParams();
                data.append("name", newName);

                fetch("player" + playerNo, {
                    method: 'post',
                    body: data,
                }).then(result => {
                    console.log(result)
                });
            } else if(obj.hasClass("ball")) {
                let player = obj.data('player')
                fetch('/game/ball/' + player).then(r => console.log(r));
            } else if(obj.hasClass('switch-mode')) {
                fetch('/game/switch-mode').then(r => console.log(r));
            } else {
                let player = obj.data('player')
                if(obj.hasClass("addPoint")) {
                    fetch('/point/inc/player' + player).then(r => console.log(r));
                } else if(obj.hasClass("decPoint")) {
                    fetch('/point/dec/player' + player).then(r => console.log(r));
                }

            }

        })
    }

    start() {
        this.start = true;
    }
}

class Player {
    constructor(no, config) {
        this.config = config
        this.no = no;
        this.name = "Player " + no
        this.points = 1;
        this.dom = $('#player' +  no);
        this.ball = false;
        this.winning = false;
    }

    create_UUID(){
        let dt = new Date().getTime();
        return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function (c) {
            let r = (dt + Math.random() * 16) % 16 | 0;
            dt = Math.floor(dt / 16);
            return (c === 'x' ? r : (r & 0x3 | 0x8)).toString(16);
        });
    }

    render() {
        if(this.config.isSetting()) {
            let name = this.dom.find('.name');

            if(name.attr("id") === undefined) {
                name.attr("id", this.create_UUID())
            }

            let ball = this.dom.find('.ball');
            ball.removeClass('btn-light');
            ball.removeClass('btn-success');
            if(this.ball) {
                ball.addClass('btn-success');
            } else {
                ball.addClass('btn-light');
            }

            if($(document.activeElement).attr('id') !== name.attr("id")) {
                name.val(this.name);
            }

            let points = this.dom.find('.points');
            if(points.attr("id") === undefined) {
                points.attr("id", this.create_UUID())
            }

            if($(document.activeElement).attr('id') !== points.attr("id")) {
                points.val(this.points);
            }
        } else {
            this.dom.removeClass('active');
            this.dom.removeClass('winning');
            if(this.winning) {
                this.dom.addClass('winning');
            } else {
                if(this.ball) {
                    this.dom.addClass('active');
                }
            }

            this.dom.find('.name').html(this.name);
            this.dom.find('.points').html(this.points);
        }
    }
}

class Game {

    constructor(config) {
        this.interval = 0;
        this.config = config;
        this.player1 = new Player(1, config);
        this.player2 = new Player(2, config);
        this.gameMode = 21;
        this.divider = $('#divider');
        this.modeEl = $('#mode');
        if(config.connection_mode === "socket") {
            this.socket = new Socket();
            this.socket.emit("game", "game started");
        }

        this.mode = config.connection_mode;
    }

    receiveGame(game) {
        this.player1.name = game.player1.name;
        this.player1.points = game.player1.points;
        this.player1.ball = game.ball === 1;
        this.player1.winning = game.winner === 1;
        this.player2.name = game.player2.name;
        this.player2.points = game.player2.points;
        this.player2.ball = game.ball === 2;
        this.player2.winning = game.winner === 2;
        this.gameMode = game.mode;
        this.player1.render();
        this.player2.render();

        this.modeEl.text(this.gameMode);
    }

    fetchGameInfo() {
        fetch("/game/info")
            .then(response => response.json())
            .then((data) => {
                    this.receiveGame(data);
                }
            )
    }

    render() {
        this.interval = 0;

        if(this.mode === "rest") {
            this.fetchGameInfo();
            this.divider.toggleClass("blue");
            this.interval = setTimeout(() => { this.render() }, 200)
        } else {
            setInterval(() => {
                if(this.divider && typeof this.divider.toggleClass === "function")
                    this.divider.toggleClass("blue");
            }, 1000);
            this.fetchGameInfo();
            this.socket.onmessage((data) => {
                try {
                    let game = JSON.parse(data)
                    this.receiveGame(game);
                } catch(e) {
                    console.error(e);
                }

            });
        }
    }

}

export { Game, Buttons, Config }