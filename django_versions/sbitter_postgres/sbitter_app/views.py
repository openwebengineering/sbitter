from django import forms
from django.contrib import messages
from django.contrib.auth import authenticate, login, logout
from django.contrib.auth.decorators import login_required
from django.contrib.auth.models import User
from django.core import serializers
from django.core.context_processors import csrf
from django.core.mail import send_mail
from django.http import HttpResponse, HttpResponseRedirect
from django.shortcuts import render_to_response, get_object_or_404, render, \
    redirect
from django.template import loader, RequestContext
from django.views.decorators.csrf import csrf_exempt

from sbitter_app.models import *
from sbitter_app.model_forms import *
from sbitter_app.forms import *

def index(request):
    return render(request, "index.html", locals())

@login_required
def view_my_sbits(request):
    user = User.objects.get(user=request.user)
    sbits = Sbit.objects.filter(user=user).order_by('-created_at')
    return render(request, 'sbits_list.html', locals())

def view_all_sbits(request):
    sbits = Sbit.objects.all().values('user', 'message', 'created_at')
    return render(request, 'sbits_list.html', locals())
