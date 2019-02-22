#functional programming
from functools import partial
from operator import add
def add1(x):
  return add(1,x)
#above function can be written as below
add1=partial(add,1) 

#lazy evaluation trick

#iterators
iter(x)
#calls next
next(x)

#this will only evaluate when it is called(in some sequence)
def lazy_integers(n=0):
  while True:
    yield n
  n+1
xs=lazy_integers()
[next(xs) for _ in range(10)]

#generator comprehensions
squares=(x**2 for x in lazy_integers())
next(squares)



#itertools

#count lazy eval
from itertools import count
count([start=0,[step=1]) #start, start+step, start+2*step, ...
#islice is a lazy eval of returning a part of a list
from itertools improt islice
islice(seq, [start=0], stop, [ste=1])
#copies a iterator
tee(it,[n=2])

#filter a list
list(filter(lambda x: x < 0, number_list))

#repeat
repeat(it,[n=forever])
#cycle = repeat forever
cycle(p)
#chain, chain up two or more lists
chain(a,b,...)
#accumulate, similar to reduce in scala
from itertools import accumulate
accumulate(p,[func=add])

#compress, filter based on true/false value in select_pattern
itertools.compress(input, select_pattern)

#dropwhile, drop elements when func return is true, it will break upon first unmatch, use this with sorted list
dropwhile(func, data)
#takewhile, take elements when func return is true, it will break upon first unmatch value, use this with sorted list
takewhile(func,data)
#groupby, group data by key, sort input before use groupby
#returns key and list
groupby(data, key = lambda x: x['key_value'])
#starmap, will iterate through data and use the operator provided to run accumulate against the unpaced data
startmap(operator.mul, data)
#tee(input), make deep copy of input
a,b = tee(input)
#zip_longest(input1,input2,fillvalue=None), this will zip input1 with input2 as index, and fill the empty element with fillvalue
#combinations(input, n), return a list of unique every n combination
#combinations_with_replacement(), same above but will add pair with itself
#product(input1,input2), similar to mysql inner join, returns cascadian product
#permutations(iterable, r=None), return r length of permutations of elements in the iterable, if r is none, it will return full-length permuataions




################TOOLZ###############################
#http://toolz.readthedocs.io/en/latest/api.html#toolz.functoolz.compose
#compose n functions
import toolz
>>> inc = lambda i: i + 1
import toolz.functoolz.compose(str, inc)(3)
'4'

#this is the same as f(g(h(x, y)))
toolz.functoolz.compose(f, g, h)(x, y)


#equivalent of h(g(f(data)))
toolz.functoolz.pipe(data, f, g, h)

#function that converts binary output to true or false
toolz.functoolz.complement(function_of_binary_output)

#converts any function to composable function
#toolz.functoolz.do(func, x)
>>> from toolz import compose
>>> from toolz.curried import do
>>> log = []
>>> inc = lambda x: x + 1
>>> inc = compose(inc, do(log.append))
>>> inc(1)
2
>>> inc(11)
12
>>> log
[1, 11]


  #manual implementation:
  #import functools
  #def compose(*functions):
  #    def compose2(f, g):
  #        return lambda x: f(g(x))
  #    #applies compose2 to each functions and initial value is lambda x:x  which is anonymous function of returning original
  #    return functools.reduce(compose2, functions, lambda x: x)
  ##the order the from right to left since it's a reduce function
  #composed_func = compose(str,lambda x: x*x, lambda x:x*2)
  #print composed_func(10)




#install specific version of dependency
pip install django_modeltranslation==0.4.0-beta2

#merge two json, or update one json with another json
c = {key: value for (key, value) in (a.items() + b.items())}
#loads json string to json
obj = json.loads(json_string) 



#python timeit example
import timeit
def strcmp():
    if "093050201114" in "093050201114000":
        return True

print timeit.timeit(stmt = 'strcmp()',setup="from __main__ import strcmp", number=100000)



#call bash script
subprocess.check_output(["echo", "Hello World!"])
'Hello World!\n'

subprocess.check_output("exit 1", shell=True)
Traceback (most recent call last):
   ...
subprocess.CalledProcessError: Command 'exit 1' returned non-zero exit status 1

#argparse -> an wrapper to parse sys.argv
import argparse
parser = argparse.ArgumentParser()
parser.add_argument("var1", help = "var1 is the first argument here", type = int) # type is defaulted to str
parser.add_argument("-rf","--random_flag", help = "--random_flag is the demo for flag argument", action = "store_true") # set -rf and --random_flag argument. will convert this args.random_flag to true if it is set
args = parser.parser_args()
print(args.var1)

#install pip inside python script
def install(package):
    pip.main(['install',package])
packages = [boto==2.48.0,
for package in packages:
    install(package)

#python3
line_arr = line.decode('cp1252').split('|')

#force reinstall pip package
pip install --upgrade --force-reinstall

#disable ssl authenticate on urllib2
ctx = ssl.create_default_context()
ctx.check_hostname = False
ctx.verify_mode = ssl.CERT_NONE

with contextlib.closing(urllib2.urlopen(url, context = ctx)) as urlret:
  query = urlret.read()



#use python3 for py.test
#make sure you are in python3 venv
python3 -m py.test ./utest_split_by_ticker.py

#check if dir exists
os.path.exists("/home/el/myfile.txt")
#get class name from a class method def foo(self)
self.__name__

#python equivalent of pwd, current directory
import os
cwd = os.getcwd()

#pdb debugging tip
#lunch another python interactive shell within pdb
!import code; code.interact(local=vars())

#sort python list by columns
>>> from operator import itemgetter
>>> L=[[0, 1, 'f'], [4, 2, 't'], [9, 4, 'afsd']]
>>> sorted(L, key=itemgetter(1,2))
[[9, 4, 'afsd'], [0, 1, 'f'], [4, 2, 't']]


#py.test show code coverage
pip install pytest-cov
py.test --cov-report term-missing --cov=weathertracker ./weathertracker/


#Python bit manipulation
#2 based 
int('00100001', 2) ==> 33 ==> 2**5 +1
#hex string
print "0x%x" % int('11111111', 2) => 0xff
#from ascii integer to char
chr(int('111011', 2)) => ';'
1 << 7 ==> 128


#unittest with mock
with patch('requests.get') as mock_requests:
    mock_requests.return_value="123"
    assert module1.method_to_test1("http://www.google.com") == "123"
#mock with decorator pattern
@patch('requests.get') # this will be the mockked first argument of the test case "mock1" 
def test_method2_wth(self, mock1): #this mock1 declared the @patch("request.get") above is named as "mock1"
    mock1.return_value = "123"
    assert wth.wth1("http://www.google.com") == "123"

#mock with side_effect
@patch('requests.get')
def test_method3_wth_with_sideeffect(self, mock1):

    def side_effect(*args, **kwargs): #here is the side effect mock, it will run when the mock method is called
        raise Exception("oops")

    mock1.side_effect = side_effect
    try:
        wth.wth1("http://www.google.com") == "123"
    except Exception as e:
        assert True


#full unit test example below
#code to mock
import requests
def wth1(url:str):
    res = requests.get(url)
    return res[:10]

#mocked unit test example below
import unittest
import wth
from mock import patch

class TestWth(unittest.TestCase):
    def test_wth(self):
        with patch('requests.get') as mock_requests:
            mock_requests.return_value="123"
            assert wth.wth1("http://www.google.com") == "123"

    @patch('requests.get')
    def test_method2_wth(self, mock1):
        mock1.return_value = "123"
        assert wth.wth1("http://www.google.com") == "123"

    @patch('requests.get')
    def test_method3_wth_with_sideeffect(self, mock1):
        def side_effect(*args, **kwargs):
            raise Exception("oops")
        mock1.side_effect = side_effect
        try:
            wth.wth1("http://www.google.com") == "123"
        except Exception as e:
            assert True
