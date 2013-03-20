from django.conf.urls.defaults import *
from django.contrib.auth.views import login, logout
from django.conf import settings

urlpatterns = patterns('sbitter_app.views',
    url(r'^$', 'index', name='index'),
    url(r'^my_sbits/$', 'view_my_sbits', name='view_my_sbits'),
    url(r'^sbits/$', 'view_all_sbits', name='view_all_sbits'),
    url(r'^login/$', login, kwargs=dict(template_name='login.html'),
        name='login'),
    url(r'^logout/$', logout, kwargs=dict(next_page='/'),
        name='logout'),
    ####
    url(r'^post_sbit/$', 'post_sbit', name='post_sbit'),
)
#Make sure _not_ to have the JSON POST Sbit view be available in production...
if settings.DEBUG:
    urlpatterns += patterns('sbitter_app.views',
    url(r'^post_sbit_json/$', 'post_sbit_json',
        name='post_sbit_json'))
