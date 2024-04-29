import {Game, Buttons, Config} from './base.mjs'

let game = new Game(new Config("setting", "rest"));
game.render();

let buttons = new Buttons();
buttons.start();