var gpio = require('pi-gpio');
var fs = require('fs');

desccribe('test pi-gpio', function(){

  mock({
    '/sys/class/gpio/gpio23/direction' : ''
  });

  gpio.open(16, 'output', function(err){

    const direction = fs.readFileSync('/sys/class/gpio/gpio23/direction').toString();

    should(direction).equal('out');

    done();
  });
});