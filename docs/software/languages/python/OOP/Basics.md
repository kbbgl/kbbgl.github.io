# Inheritance

```python

class Pet:
    def __init__(self, name, age):
        self.name = name
        self.age = age
    

class Cat(Pet):
    def __init__(self, name, age, color):
        super().__init__(age, name)
        self.color = color
        
```

We can use class method decorator to call a method of a class that is not specifically associated to any instance of a class:

```python

class Person():
    number_of_people = 0
    
    def __init__(self, name):
        self.name = name
        
    @classmethod
    def number_of_people_(cls):
        return cls.number_of_people
    
    @classmethod
    def add_person(cls):
        cls.number_of_people += 1
        
p1 = Person("Time")
p2 = Person("Jill")
print(Person.number_of_people_)

# Will print 2

```

## Static functions

```python
class Math:
    
    @staticmethod
    def add5(x):
        return x + 5


print(Math.add5(10))
```
