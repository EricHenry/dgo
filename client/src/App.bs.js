'use strict';

var React = require("react");
var Caml_option = require("bs-platform/lib/js/caml_option.js");
var Navbar$Client = require("./Navbar.bs.js");
var MaterialUi_Box = require("@jsiebern/bs-material-ui/src/MaterialUi_Box.bs.js");
var MaterialUi_Grid = require("@jsiebern/bs-material-ui/src/MaterialUi_Grid.bs.js");
var Core = require("@material-ui/core");
var MaterialUi_Container = require("@jsiebern/bs-material-ui/src/MaterialUi_Container.bs.js");
var MaterialUi_Typography = require("@jsiebern/bs-material-ui/src/MaterialUi_Typography.bs.js");
var MaterialUi_CssBaseline = require("@jsiebern/bs-material-ui/src/MaterialUi_CssBaseline.bs.js");

var gameDescription = "This is a description of the game";

function App(Props) {
  return React.createElement(React.Fragment, undefined, React.createElement(Core.CssBaseline, MaterialUi_CssBaseline.makeProps(undefined, undefined, undefined, undefined, undefined, undefined, undefined)), React.createElement(Navbar$Client.make, {}), React.createElement(Core.Box, MaterialUi_Box.makeProps(undefined, undefined, Caml_option.some(React.createElement(Core.Container, MaterialUi_Container.makeProps(Caml_option.some(React.createElement(Core.Grid, MaterialUi_Grid.makeProps(undefined, /* Center */980392437, Caml_option.some(null), undefined, undefined, true, /* Column */-81804554, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined), React.createElement(Core.Grid, MaterialUi_Grid.makeProps(undefined, undefined, Caml_option.some(React.createElement(Core.Typography, MaterialUi_Typography.makeProps(undefined, gameDescription, undefined, undefined, undefined, undefined, undefined, undefined, undefined, /* H4 */16108, undefined, undefined, undefined, undefined, undefined, undefined, undefined))), undefined, undefined, undefined, undefined, true, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined)), React.createElement(Core.Grid, MaterialUi_Grid.makeProps(undefined, undefined, "Another Grid item", undefined, undefined, undefined, undefined, true, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined)))), undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined))), {
                      HASH: /* String */-976970511,
                      VAL: "span"
                    }, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined)));
}

var make = App;

exports.gameDescription = gameDescription;
exports.make = make;
/* react Not a pure module */
