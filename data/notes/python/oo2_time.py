#!/usr/bin/python
# Think Python, Ch. 17 Classes and methods

class Time():
    """Represents the time of day.
    attributes: hour, minute, second"""

    # init method gets invoked when an object is instantiated
    def __init__(self, hour=0, minute=0, second=0):
        self.hour = hour
        self.minute = minute
        self.second = second

    # gets invoked when you print an object (useful for debugging)
    def __str__(self):
        return "%.2d:%.2d:%.2d" % (self.hour, self.minute, self.second)

    def print_time(self):
        print "%.2d:%.2d:%.2d" % (self.hour, self.minute, self.second)

    def time_to_int(self):
        minutes = self.hour * 60 + self.minute
        seconds = minutes * 60 + self.second
        return seconds

    def increment(self, seconds):
        seconds += self.time_to_int()
        return int_to_time( seconds )

    def is_after(self, other):
        return self.time_to_int() > other.time_to_int()

def int_to_time(seconds):
    t = Time()
    minutes, t.second = divmod(seconds, 60)
    t.hour, t.minute = divmod(minutes, 60)
    return t

def print_attributes(obj):
    for attr in vars(obj):
        print(attr, getattr(obj, attr))

time0 = Time()
print time0
print_attributes(time0)

time1 = Time()
time1.hour = 7
time1.minute = 25
time1.second = 0

time2 = Time()
time2.hour = 7
time2.minute = 40
time2.second = 0

time2 = time2.increment(60)

for t in time0, time1, time2:
    t.print_time()

if time2.is_after(time1):
    print 'time2 is after time1'
