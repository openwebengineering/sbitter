from django.conf.urls.defaults import *
from django.contrib.auth.views import login, logout

urlpatterns = patterns('sbitter_app.views',
    url(r'^$', 'index', name='index'),
    url(r'^my_sbits/$', 'view_my_sbits', name='view_my_sbits'),
    url(r'^sbits/$', 'view_all_sbits', name='view_all_sbits'),
    url(r'^login/$', login, kwargs=dict(template_name='login.html'),
        name='login'),
    url(r'^logout/$', logout, kwargs=dict(next_page='/'),
        name='logout'),
)
