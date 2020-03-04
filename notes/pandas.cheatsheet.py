# View preTestscore where postTestscore is greater than 50
df['preTestScore'].where(df['postTestScore'] > 50) 

#get pandas columns headers
f.columns.values


#list first 5 elements of pandas framework
f.head()

#select columns from df
f[['column1_name','name3']]

#select columns from index
f.ix[:,0:2]

# select a range of df
df[8:9]

#load csv file
f = pandas.read_csv('some_file.gz', compression = 'gzip',delimiter= '|')


#lookup value and add that as a new column
#df.lookup(df_to_be_updated_index, desired_value_series)
df1,df2

df2.lookup(df1.rowlabel, df1.columnlabel)

#another example
f.lookup(null_df.index,null_df['name5'])

#selecting columns by label
df.loc[:,['A','B']]
#select by position
df.iloc[3]


#select by condition
df[df.A > 0]


#filtering
df2[df2['E'].isin(['two','four'])]

# string pattern match
df[df['ids'].str.contains("ball")]

#set value by label
df.at[dates[0],'A'] = 0 

#set value by position
df.iat[0,1] = 0 


#reverse all positive numbers in df2 into negative numbers
df2[df2 > 0] = -df2

#pandas null/empty value lookup
f2[f2.column2_name.isnull()].head()


#pandas select by regex
f2[f2.symbol.str.contains('\.')].head()

#is not in
null_df[~null_df.name5.isin(f.column3_name)]

#grab a column from another df and join to original df as a new column
  #have to index both data column so they can join on same index
new_f = f[['column1_name','column2_name']].set_index('name4')
new_null_df = null_df.set_index('name5')
new_null_df.update(new_f)


#delete a column
del df['column_name']
#move index as a new column
#either:

df['index1'] = df.index
#or, .reset_index:
df.reset_index(level=0, inplace=True)


# merge df
pd.merge(left, right, how='inner', on=None, left_on=None, right_on=None,
         left_index=False, right_index=False, sort=True,
         suffixes=('_x', '_y'), copy=True, indicator=False,
         validate=None)


# change a column values
df['col'] = 'str' + df['col'].astype(str)
# map function for pandas
data['result'] = data['result'].map(lambda x: x.lstrip('+-').rstrip('aAbBcC'))

# filter a pd df with another pd df
df1 = pd.DataFrame({'c': ['A', 'A', 'B', 'C', 'C'],
                    'k': [1, 2, 2, 2, 2],
                    'l': ['a', 'b', 'a', 'a', 'd']})
df2 = pd.DataFrame({'c': ['A', 'C'],
                    'l': ['b', 'a']})
keys = list(df2.columns.values)
i1 = df1.set_index(keys).index
i2 = df2.set_index(keys).index
df1[~i1.isin(i2)]


# inner join(filter rows based on another pdf column value) two pandas df
# You could use a simple mask:
m = df2.SKU.isin(df1.SKU)
df2 = df2[m]
# or
# You are looking for an inner join. Try this:
df3 = df1.merge(df2, on=['SKU','Sales'], how='inner')

#  SKU  Sales
#0   A    100
#1   B    200
#2   C    300

# Or this:

df3 = df1.merge(df2, on='SKU', how='inner')

#  SKU  Sales_x  Sales_y
#0   A      100      100
#1   B      200      200
#2   C      300      300

# filter pandas row based on column value using regex
df.b.str.contains('^f')  # this returns true

# not true using regex
true_arr = current_mislabelled_segment_df.Token.str.contains('AA_')
new_df_false = current_mislabelled_segment_df[~true_arr]

df.drop_duplicates('COL2')
# transpose would change pandas dataframe from vertical print to horizontal print
df.T
# 0    1    2
# 0  100  200  300

# sum a column
pd.Series([]).sum()


# group by
df.groupby("committee_name_x").amount.sum()

# Sorting the biggest totals to the top is as easy as appending the sort_values trick we already know to the end. And voila there’s our answer.
merged.groupby("committee_name_x").amount.sum().reset_index().sort_values("amount", ascending=False)

# select few columns from a df
apts[['homeType', 'price', 'livingArea']]

# drops na when there are two or more nan in pandas
nms.dropna(thresh=2)

# filters a df with one of its column not empty
inms[nms.name.notnull()]

# pandas filter rows on multiple condition, with example of "and" and "or"
apts[apts.price.notnull() & apts.livingArea.notnull() | False]


