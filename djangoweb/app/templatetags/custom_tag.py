import random
from django import template

register = template.Library()

register.filter("random_title", random_title)

@register.filter(name="random_title")
def random_title(value):
    title = ["Naruto", "Detective Conan", "Bleach"]
    return value + random.choice(title)
