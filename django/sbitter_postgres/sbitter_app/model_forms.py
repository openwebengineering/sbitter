from django.db import models
from django.contrib.auth.models import User
from django.contrib import admin
from django import forms

from tagging.fields import TagField
from sbitter_app.models import *

import datetime
import os

class SbitForm(forms.ModelForm):
    message = forms.CharField(label="Message")

    class Meta:
        model = Sbit
        fields = ("message",)
