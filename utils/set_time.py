#
# Set time on charge controllers based on local time while running this script
# comment/uncomment write commands below to view or set time
#
from pyepsolartracer.client import EPsolarTracerClient
from pyepsolartracer.registers import registers,coils,registerByName
from pymodbus.client.sync import ModbusSerialClient

ports = [ '/dev/ttyXRUSB0', '/dev/ttyXRUSB1' ]

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

    #register = registerByName("Real time clock 1")
    #print vars(register)

    addr = 0x9013
    vals = [ minuteAndSecond, dayAndHour, yearAndMonth]

    response = client.read_input("Real time clock 1")
    print "Seconds: " + str(int(response) & 0xFF)
    print "Minutes: " + str(int(response) >> 8)

    response = client.read_input("Real time clock 2")
    print "Hours: " + str(int(response) & 0xFF)
    print "Day: " + str(int(response) >> 8)

    response = client.read_input("Real time clock 3")
    print "Month: " + str(int(response) & 0xFF)
    print "Year: " + str(int(response) >> 8)

    #response = client.client.write_registers(addr, vals, unit = client.unit)
    #print "Write result: " + str(vars(response))
    client.close()
