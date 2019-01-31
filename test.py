#!/usr/bin/env python3
#-*- coding:utf-8 -*-
import mitmproxy


def request(flow):
	if flow.request.method == 'POST':
		if flow.request.host == 'kslabs.ru': 
			if flow.request.path == '/service/19032014/do.php':
				#flow.request.url == http://kslabs.ru/service/19032014/do.php
				if len(flow.request.content)>0 :
					flow.request.content = b'serial=155B-38B1-0EC7-2E29&deviceID=9121369032a98e208481689c178e47d6&emails=yunhmong@gmail.com &version=v3.62'
	
