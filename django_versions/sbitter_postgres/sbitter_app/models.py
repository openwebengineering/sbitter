from django.db import models
from django.contrib.auth.models import User
from django.contrib import admin
from django import forms

from tagging.fields import TagField

import datetime
import os

class Base(models.Model):
    created_at  = models.DateTimeField(auto_now=True)
    modified_at = models.DateTimeField(auto_now=True, auto_now_add=True)

    class Meta:
        abstract = True

class Sbit(Base):
    user    = models.ForeignKey(User)
    message = models.CharField(max_length=140)

    def __unicode__(self):
        return self.user.username + " sayed " + self.message
