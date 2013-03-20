"""
This file demonstrates writing tests using the unittest module. These will pass
when you run "manage.py test".

Replace this with more appropriate tests for your application.
"""

from django.test import TestCase
from django.test.client import Client

class SimpleTest(TestCase):
    def test_basic_addition(self):
        """
        Tests that 1 + 1 always equals 2.
        """
        self.assertEqual(1 + 1, 2)


class SbitGETTestCase(TestCase):
    def setUp(self):
        global c
        c = Client()

    def test_get(self):
        request = c.get('/')
        self.assertEqual(request.status_code, 200)

    def test_getting_sbits(self):
        request = c.get('/sbits/')
        self.assertEqual(request.status_code, 200)
