var gpio = require('gpio');

var gpio4 = gpio.export(4, {
    direction:'out',
    interval :200,
    ready : function(){
        console.log('ok');
    }
});

console.log(gpio4);

gpio4.set();

console.log("hello world");
