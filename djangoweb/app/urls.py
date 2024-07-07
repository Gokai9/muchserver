from django.urls import path
from app.views import Hello, HelloTemplate

urlpatterns = [
    path('', Hello, name="hello"),
    path('posts', HelloTemplate)
]