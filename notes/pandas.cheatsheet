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
either:

df['index1'] = df.index
or, .reset_index:

df.reset_index(level=0, inplace=True)

