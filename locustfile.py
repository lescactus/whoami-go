#!/usr/bin/python

import random
from locust import HttpUser, TaskSet, between

def index(l):
    l.client.get("/")

def ip(l):
    l.client.get("/ip")

def port(l):
    l.client.get("/port")

def lang(l):
    l.client.get("/lang")

def ua(l):
    l.client.get("/ua")

def rawGo(l):
    l.client.get("/raw/go")

def rawJSON(l):
    l.client.get("/raw/json")

def rawYAML(l):
    l.client.get("/raw/yaml")

class UserBehavior(TaskSet):
    def on_start(self):
        index(self)

    tasks = {index: 2,
        ip: 10,
        port: 5,
        lang: 9,
        ua: 4,
        rawGo: 110,
        rawJSON: 6,
        rawYAML: 30,
        }
    
class WebsiteUser(HttpUser):
    tasks = [UserBehavior]
    wait_time = between(1, 2)


