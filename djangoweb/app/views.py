from django.http import HttpResponse, HttpRequest
from django.shortcuts import render
from requests import get

# Create your views here.
def Hello(request: HttpRequest) -> HttpResponse:
    response_content = """
    <html>
    <head><title>Hello World!</title>
    <body>
        <h1>This is a demo.</h1>
    </body>
    </html>
    """
    return HttpResponse(response_content)

def HelloTemplate(request: HttpRequest):
    res = get('https://jsonplaceholder.typicode.com/posts')
    posts = res.json()
    data = {"posts": posts}
    return render(request, "index.html", data)