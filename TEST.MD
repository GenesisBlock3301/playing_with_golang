# Automated-testing tool

[![build-status-image]][build-status]
[![coverage-status-image]][codecov]
[![pypi-version]][pypi]

**Awesome web-browsable Web APIs.**

Full documentation for the project is available at [https://demo-documentation.com/][docs].

---


# Sponsored By Brain Station-23

Automated-testing tool is a *Brain station's product*. If you use
this tool commercially we strongly encourage you to invest in its
continued development by [signing up for a paid plan][funding].

The initial aim is to provide a single full-time position on REST framework.
*Every single sign-up makes a significant impact towards making that possible.*

<a href="https://brainstation-23.com/?bc"><img align="center" width="300" height="300" src="https://media-exp1.licdn.com/dms/image/C510BAQEhBWM1jX08fw/company-logo_200_200/0/1566279758471?e=1671667200&v=beta&t=94o91mt8ai0zvqMRZyhXvv5psTQTVkqMe10NREJ9rQg"></a>

[//]: # ([![][bs-image]][bs-url])


# Overview

Automated-testing tool is a powerful and flexible toolkit for [SQA][sqa] or [software][software] engineers.
This is a python based Web Automation testing tool with
dynamic Design Pattern adaptability and a full stack solution.

Some reasons you might want to use this tool:

* This is a testing supporting tool for software and SQA engineer that will reduce redundancy of testing phase.
* By default, this application has [headless mode][headless-mode] (can be changed). It will generate automated QA report and sent through [SMTP][smtp] email to the end users.
* Complete solution for both [Black-box][black-box] and [White-box][white-box] testing.
* Robot feature which is a [Terminal/CMD ][terminal-cmd]based tool will bring much more mobility and efficiency for Automation Engineer to design and perform any Automation Test.
*  The biggest advantage is that this framework by default follows
   [POM design pattern][pom] and tried to bring combination from [TestNG][test-ng] & [Robot Framework][robot-framework].


[//]: # (There is a live example API for testing purposes, [available here][sandbox].)

**Below**: *Screenshot from the browsable API*

[//]: # (![Screenshot][image])

----

# Requirements

* Python (3.10+)
* Django
* Pep8-naming
* Pytest-factoryboy
* Faker
* Django-liveserver
* Selenium
* Cryptography
* Webdriver-manager
* Pytest-django
* Openpyxl
* Punq
* Beautifulsoup4
* Lxml
* Pytest-html
* Pytest-json-report

We **highly recommend** and only officially support the latest patch release of
each Python and Django series.

# Installation

Install `pipenv` using `pip`

    pip install pipenv


Setup your environment using `pipenv` ...

    pipenv shell

Add all dependencies to your environment by using:

    pipenv install


# Example

Let's take a look at a quick example of using REST framework to build a simple model-backed API for accessing users and groups.



Now edit the `example/urls.py` module in your project:

```python
from django.contrib.auth.models import User
from django.urls import include, path
from rest_framework import routers, serializers, viewsets


# Serializers define the API representation.
class UserSerializer(serializers.HyperlinkedModelSerializer):
    class Meta:
        model = User
        fields = ['url', 'username', 'email', 'is_staff']


# ViewSets define the view behavior.
class UserViewSet(viewsets.ModelViewSet):
    queryset = User.objects.all()
    serializer_class = UserSerializer


# Routers provide a way of automatically determining the URL conf.
router = routers.DefaultRouter()
router.register(r'users', UserViewSet)

# Wire up our API using automatic URL routing.
# Additionally, we include login URLs for the browsable API.
urlpatterns = [
    path('', include(router.urls)),
    path('api-auth/', include('rest_framework.urls', namespace='rest_framework')),
]
```

We'd also like to configure a couple of settings for our API.

Add the following to your `settings.py` module:

```python
INSTALLED_APPS = [
    ...  # Make sure to include the default installed apps here.
    'rest_framework',
]

REST_FRAMEWORK = {
    # Use Django's standard `django.contrib.auth` permissions,
    # or allow read-only access for unauthenticated users.
    'DEFAULT_PERMISSION_CLASSES': [
        'rest_framework.permissions.DjangoModelPermissionsOrAnonReadOnly',
    ]
}
```

That's it, we're done!

    ./manage.py runserver

You can now open the API in your browser at `http://127.0.0.1:8000/`, and view your new 'users' API. If you use the `Login` control in the top right corner you'll also be able to add, create and delete users from the system.

You can also interact with the API using command line tools such as [`curl`](https://curl.haxx.se/). For example, to list the users endpoint:

    $ curl -H 'Accept: application/json; indent=4' -u admin:password http://127.0.0.1:8000/users/
    [
        {
            "url": "http://127.0.0.1:8000/users/1/",
            "username": "admin",
            "email": "admin@example.com",
            "is_staff": true,
        }
    ]

Or to create a new user:

    $ curl -X POST -d username=new -d email=new@example.com -d is_staff=false -H 'Accept: application/json; indent=4' -u admin:password http://127.0.0.1:8000/users/
    {
        "url": "http://127.0.0.1:8000/users/2/",
        "username": "new",
        "email": "new@example.com",
        "is_staff": false,
    }

# Documentation & Support

Full documentation for the project is available at [https://www.django-rest-framework.org/][docs].

For questions and support, use the [REST framework discussion group][group], or `#restframework` on libera.chat IRC.

You may also want to [follow the author on Twitter][twitter].

# Security

Please see the [security policy][security-policy].

[headless-mode]:https://en.wikipedia.org/wiki/Headless_browser#:~:text=A%20headless%20browser%20is%20a,interface%20or%20using%20network%20communication.
[black-box]: https://en.wikipedia.org/wiki/Black-box_testing
[white-box]: https://en.wikipedia.org/wiki/White-box_testing
[terminal-cmd]: https://en.wikipedia.org/wiki/Command-line_interface
[smtp]: https://en.wikipedia.org/wiki/Simple_Mail_Transfer_Protocol
[sqa]: https://en.wikipedia.org/wiki/Software_quality_assurance/
[software]: https://en.wikipedia.org/wiki/Software
[pom]: https://medium.com/tech-tajawal/page-object-model-pom-design-pattern-f9588630800b
[test-ng]: https://testng.org/doc/
[robot-framework]: https://robotframework.org/SeleniumLibrary/SeleniumLibrary.html
[bs-image]: https://media-exp1.licdn.com/dms/image/C510BAQEhBWM1jX08fw/company-logo_200_200/0/1566279758471?e=1671667200&v=beta&t=94o91mt8ai0zvqMRZyhXvv5psTQTVkqMe10NREJ9rQg
[bs-url]: https://brainstation-23.com/?bc




[build-status-image]: https://github.com/encode/django-rest-framework/actions/workflows/main.yml/badge.svg
[build-status]: https://github.com/encode/django-rest-framework/actions/workflows/main.yml
[coverage-status-image]: https://img.shields.io/codecov/c/github/encode/django-rest-framework/master.svg
[codecov]: https://codecov.io/github/encode/django-rest-framework?branch=master
[pypi-version]: https://img.shields.io/pypi/v/djangorestframework.svg
[pypi]: https://pypi.org/project/djangorestframework/
[twitter]: https://twitter.com/starletdreaming
[group]: https://groups.google.com/forum/?fromgroups#!forum/django-rest-framework
[sandbox]: https://restframework.herokuapp.com/

[funding]: https://fund.django-rest-framework.org/topics/funding/
[sponsors]: https://fund.django-rest-framework.org/topics/funding/#our-sponsors

[stream-img]: https://raw.githubusercontent.com/encode/django-rest-framework/master/docs/img/premium/stream-readme.png
[spacinov-img]: https://raw.githubusercontent.com/encode/django-rest-framework/master/docs/img/premium/spacinov-readme.png
[retool-img]: https://raw.githubusercontent.com/encode/django-rest-framework/master/docs/img/premium/retool-readme.png
[bitio-img]: https://raw.githubusercontent.com/encode/django-rest-framework/master/docs/img/premium/bitio-readme.png
[posthog-img]: https://raw.githubusercontent.com/encode/django-rest-framework/master/docs/img/premium/posthog-readme.png
[cryptapi-img]: https://raw.githubusercontent.com/encode/django-rest-framework/master/docs/img/premium/cryptapi-readme.png
[fezto-img]: https://raw.githubusercontent.com/encode/django-rest-framework/master/docs/img/premium/fezto-readme.png

[sentry-url]: https://brainstation-23.com/?bc
[stream-url]: https://getstream.io/?utm_source=DjangoRESTFramework&utm_medium=Webpage_Logo_Ad&utm_content=Developer&utm_campaign=DjangoRESTFramework_Jan2022_HomePage
[spacinov-url]: https://www.spacinov.com/
[retool-url]: https://retool.com/?utm_source=djangorest&utm_medium=sponsorship
[bitio-url]: https://bit.io/jobs?utm_source=DRF&utm_medium=sponsor&utm_campaign=DRF_sponsorship
[posthog-url]: https://posthog.com?utm_source=drf&utm_medium=sponsorship&utm_campaign=open-source-sponsorship
[cryptapi-url]: https://cryptapi.io
[fezto-url]: https://www.fezto.xyz/?utm_source=DjangoRESTFramework

[oauth1-section]: https://www.django-rest-framework.org/api-guide/authentication/#django-rest-framework-oauth
[oauth2-section]: https://www.django-rest-framework.org/api-guide/authentication/#django-oauth-toolkit
[serializer-section]: https://www.django-rest-framework.org/api-guide/serializers/#serializers
[modelserializer-section]: https://www.django-rest-framework.org/api-guide/serializers/#modelserializer
[functionview-section]: https://www.django-rest-framework.org/api-guide/views/#function-based-views
[generic-views]: https://www.django-rest-framework.org/api-guide/generic-views/
[viewsets]: https://www.django-rest-framework.org/api-guide/viewsets/
[routers]: https://www.django-rest-framework.org/api-guide/routers/
[serializers]: https://www.django-rest-framework.org/api-guide/serializers/
[authentication]: https://www.django-rest-framework.org/api-guide/authentication/
[image]: https://www.django-rest-framework.org/img/quickstart.png

[docs]: https://www.django-rest-framework.org/
[security-policy]: https://github.com/encode/django-rest-framework/security/policy