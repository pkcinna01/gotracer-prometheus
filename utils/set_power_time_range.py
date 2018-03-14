#
# Set time when charge controllers enable load terminals
#
# There are two time spans that can be set.  For example one
# can be set to define when fans are enabled in afternoon and 
# the other could be for turning on security lights at night
#
# comment/uncomment write commands below to view or set
#

from pyepsolartracer.client import EPsolarTracerClient
from pyepsolartracer.registers import registers,coils,registerByName
from pymodbus.client.sync import ModbusSerialClient
from sys import stdout

#ports = [ '/dev/ttyXRUSB0', '/dev/ttyXRUSB1' ]
ports = [ '/dev/ttyXRUSB1' ]

for port in ports:
    print "************ " + port + " ***************"
    serialclient = ModbusSerialClient(method = 'rtu', port = port, baudrate = 115200)
    client = EPsolarTracerClient(serialclient = serialclient)
    client.connect()

    import datetime
    now = datetime.datetime.now()

    print now.year, now.month, now.day, now.hour, now.minute, now.second

    yearAndMonth = int(((now.year-2000) << 8) | now.month)
    dayAndHour = int(((now.day)<<8) | now.hour)
    minuteAndSecond = int(((now.minute)<<8) | now.second)

    response = client.read_input("Real time clock 2")
    stdout.write("Device Time: " + str(int(response) & 0xFF))
    response = client.read_input("Real time clock 1")
    stdout.write(":" + str(int(response) >> 8) + ":" + str(int(response) & 0xFF))
    stdout.flush()

    print ""
    print "******************************************************"
    loadCtrlModesAddr = 0x903D
    
    # 0000H Manual Control
    # 0001H Light ON/OFF
    # 0002H Light ON+ Timer/
    # 0003H Time Control
    loadCtrlMode = 0x0003
    
    #response = client.client.write_registers(loadCtrlModesAddr, loadCtrlMode, unit = client.unit)
    #print "Write result: " + str(vars(response))
    
    response = client.read_input("Load controling modes")
    print "Load controlling mode: " + hex(response.value)

    #response = client.write_output("Turn on timing 1 hour",9)
    #response = client.write_output("Turn on timing 1 min",30)
    response = client.write_output("Turn off timing 1 hour",16)
    response = client.write_output("Turn off timing 1 min",0)
    #response = client.write_output("Turn on timing 2 hour",9)
    #response = client.write_output("Turn on timing 2 min",30)
    response = client.write_output("Turn off timing 2 hour",16)
    response = client.write_output("Turn off timing 2 min",0)
    
    print "Load Time Contolled #1"
    second = client.read_input("Turn on timing 1 sec").value
    minute = client.read_input("Turn on timing 1 min").value
    hour = client.read_input("Turn on timing 1 hour").value
    print "  ON:  %02d:%02d:%02d" % (int(hour),int(minute),int(second))
    second = client.read_input("Turn off timing 1 sec").value
    minute = client.read_input("Turn off timing 1 min").value
    hour = client.read_input("Turn off timing 1 hour").value
    print "  OFF: %02d:%02d:%02d" % (int(hour),int(minute),int(second))

    print "Load Time Contolled #2"
    second = client.read_input("Turn on timing 2 sec").value
    minute = client.read_input("Turn on timing 2 min").value
    hour = client.read_input("Turn on timing 2 hour").value
    print "  ON:  %02d:%02d:%02d" % (int(hour),int(minute),int(second))
    second = client.read_input("Turn off timing 2 sec").value
    minute = client.read_input("Turn off timing 2 min").value
    hour = client.read_input("Turn off timing 2 hour").value
    print "  OFF: %02d:%02d:%02d" % (int(hour),int(minute),int(second))

    client.close()
