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
    return render(request, "public/index.html", locals())

@login_required
def view_my_sbits(request):
    user = User.objects.get(user=request.user)
    sbits = Sbit.objects.filter(user=user).order_by('-created_at')
    return render(request, 'sbitter/sbits_list.html', locals())

def view_all_sbits(request):
    sbits = Sbit.objects.all().values('user', 'message', 'created_at')
    return render(request, 'sbitter/sbits_list.html', locals())

@login_required
def post_sbit(request):
    if request.methond == 'POST':
        form = SbitForm(request.POST)
        if form.is_valid():
            sbit = form.save(commit=False)
            sbit.user = request.user
            sbit.save()
            messages.success(request, 'New Sbit Added')
            return redirect('index')
        else:
            for field in form:
                for error in field.errors:
                    pass
    else:
        form = SbitForm()
    return render(request, 'sbitter/post_sbit.html', {'form': form})
