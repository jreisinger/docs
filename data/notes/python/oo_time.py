#!/usr/bin/python
# Think Python, Ch. 16 Classes and functions

class Time():
    """Represents the time of day.

    attributes: hour, minute, second"""

def print_time(t):
    print "%.2d:%.2d:%.2d" % (t.hour, t.minute, t. second)

def is_after(t1, t2):
    return (t1.hour, t1.minute, t1.second) > (t2.hour, t2.minute, t2.second)

# Prototype and patch
def add_time(t1, t2):
    time = Time()
    time.hour = t1.hour + t2.hour
    time.minute = t1.minute + t2.minute
    time.second = t1.second + t2.second

    if time.second >= 60:
        time.second -= 60
        time.minute += 1

    if time.minute >= 60:
        time.minute -= 60
        time.hour   += 1

    return time

# Designed development - involves high-level insight into the problem (time is actually a 60 base number!) and more planning
def increment(t, seconds):
    return int_to_time( time_to_int(t) + seconds )

def time_to_int(t):
    minutes = t.hour * 60 + t.minute
    seconds = minutes * 60 + t.second
    return seconds

def int_to_time(seconds):
    t = Time()
    minutes, t.second = divmod(seconds, 60)
    t.hour, t.minute = divmod(minutes, 60)
    return t

time1 = Time()
time1.hour = 7
time1.minute = 25
time1.second = 0

time2 = Time()
time2.hour = 7
time2.minute = 40
time2.second = 0

print_time(time1)
print_time( increment(time1, 3600) )
print_time(time1)
