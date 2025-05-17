import csv
import logging

class ReadData:

    def __init__(self):
        self.data = []
        self.logger = logging.basicConfig()
        
    def read_file(self, loc: str):
        with open(loc, newline='') as f:
            for row in f:
                self.data.append(row)
    
    def flust(self):
        if len(self.data) > 0:
            self.data = []

