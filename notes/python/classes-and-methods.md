```python3
#!/usr/bin/python3
# Think Python, Ch. 17 Classes and methods

class Time():
    """Represents the time of day.
    (data) attributes: hour, minute, second"""

    # init method gets invoked when an object is instantiated
    def __init__(self, hour=0, minute=0, second=0):
        self.hour = hour
        self.minute = minute
        self.second = second

    # gets invoked when you print an object (useful for debugging)
    def __str__(self):
        return "%.2d:%.2d:%.2d" % (self.hour, self.minute, self.second)

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
time1 = Time(7, 25, 0)
time2 = Time(7, 40, 3)

for t in time0, time1, time2:
    print(t)

time2 = time2.increment(60)
print_attributes(time2)

if time2.is_after(time1):
    print('time2 is after time1')
```

More: [ MIT OpenCourseWare - 8. Object Oriented Programming](https://www.youtube.com/watch?v=-DP1i2ZU9gk&list=PLUl4u3cNGP63WbdFxL8giv4yhgdMGaZNA&index=27)
