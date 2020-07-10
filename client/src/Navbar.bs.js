'use strict';

var React = require("react");
var Caml_option = require("bs-platform/lib/js/caml_option.js");
var Core = require("@material-ui/core");
var MaterialUi_AppBar = require("@jsiebern/bs-material-ui/src/MaterialUi_AppBar.bs.js");
var MaterialUi_Toolbar = require("@jsiebern/bs-material-ui/src/MaterialUi_Toolbar.bs.js");
var MaterialUi_Typography = require("@jsiebern/bs-material-ui/src/MaterialUi_Typography.bs.js");

function Navbar(Props) {
  return React.createElement(Core.AppBar, MaterialUi_AppBar.makeProps(undefined, undefined, undefined, undefined, Caml_option.some(React.createElement(Core.Toolbar, MaterialUi_Toolbar.makeProps(Caml_option.some(React.createElement(Core.Typography, MaterialUi_Typography.makeProps(undefined, "Don't Go Over", undefined, undefined, undefined, undefined, undefined, undefined, undefined, /* H3 */16107, undefined, undefined, undefined, undefined, undefined, undefined, undefined))), undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined))), undefined, undefined, /* Static */982536398, undefined, undefined, undefined, undefined, undefined, undefined));
}

var make = Navbar;

exports.make = make;
/* react Not a pure module */
