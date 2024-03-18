package com.btc.wallect.db;

import android.content.ContentValues;
import android.content.Context;
import android.database.AbstractWindowedCursor;
import android.database.Cursor;
import android.database.CursorWindow;
import android.database.sqlite.SQLiteDatabase;
import android.os.Build;


import androidx.annotation.RequiresApi;

import com.btc.wallect.model.entity.Wallet;

import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Date;
import java.util.List;

public class DBdao {
    private DBOpenHelper dBhelpUtil;


    /**相当于获得一个链接数据库的对象*/
    private SQLiteDatabase DB;
    private Context context;
    public DBdao(Context context, DBOpenHelper dBhelpUtil){
        this.context =context;
        this.dBhelpUtil = dBhelpUtil;
    }
    //保存数据
    public Long save(Wallet wallet) {
        /** 获取一个写 操作数据的对象*/
        DB = dBhelpUtil.getWritableDatabase();

        ContentValues contentValues = new ContentValues();
        contentValues.put(DBOpenHelper.TB_NAME,wallet.name);
        contentValues.put(DBOpenHelper.TB_SEX,wallet.password);
        contentValues.put(DBOpenHelper.TB_AGE,wallet.txt);
        contentValues.put(DBOpenHelper.TB_CLAZZ,wallet.collect);

//        Log.e("TAG","--------------"+student.toString());
//        Toast.makeText(context,"sql 语句--"+student.toString(),Toast.LENGTH_LONG).show();
        //时间
        Date date = new Date();
        //格式化
        SimpleDateFormat simpleDateFormat = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");
        contentValues.put(DBOpenHelper.TB_CREATEDATE, simpleDateFormat.format(date));


        /**insert()
         * String table: 表名
         * String nullColumnHack： 不允许插入空行，为了防止插入空行，可以在这里随便指定一列， 如果有空值插入 会用null表示，好像没作用~
         * ContentValues values 数据行数据
         * 返回值 成功插入行号的id  ,插入失败 -1
         */
        return DB.insert(DBOpenHelper.TABLE_NAME,"空值",contentValues);
        //INSERT INTO tb_student(id,age,sex,name,clazz,createDate) VALUES (?,?,?,?,?,?)

    }

    /**查询数据*/
    public List<Wallet> select(Long id) {
        /** 获取一个读 操作数据的对象*/
        DB =dBhelpUtil.getReadableDatabase();

        /**query() 查询数据
         *String table, 表名
         * String[] columns, 要查询要显示的列
         * String selection,   查询条件
         * String[] selectionArgs, 参数值
         * String groupBy, 分组
         * String having, 分组后的条件
         * String orderBy 排序
         * 返回游标 Cursor
         */
        String[] columns = new String[]{
                "id",
                DBOpenHelper.TB_NAME,
                DBOpenHelper.TB_SEX,
                DBOpenHelper.TB_AGE,
                DBOpenHelper.TB_CLAZZ,
                DBOpenHelper.TB_CREATEDATE
        };
        Cursor cursor = null;
        if(id == null){
            //全查
            cursor = DB.query(DBOpenHelper.TABLE_NAME,columns,null,null,null,null,"id desc");
        }else {
            //根据id 查询
            cursor = DB.query(DBOpenHelper.TABLE_NAME,columns,"id=?",new String[]{String.valueOf(id)},null,null,null);

        }

        List<Wallet> studentList = new ArrayList<>();
        if(cursor != null){
            //遍历游标
            while(cursor.moveToNext()){
                Wallet student = new Wallet();
                // 根据游标找到列  在获取数据
                student.id = cursor.getLong(cursor.getColumnIndexOrThrow("id"));
                student.name = cursor.getString(cursor.getColumnIndexOrThrow(DBOpenHelper.TB_NAME));
                student.password = cursor.getString(cursor.getColumnIndexOrThrow(DBOpenHelper.TB_SEX));
                student.txt = cursor.getString(cursor.getColumnIndexOrThrow(DBOpenHelper.TB_AGE));
                student.collect = cursor.getString(cursor.getColumnIndexOrThrow(DBOpenHelper.TB_CLAZZ));
                student.creatDate = cursor.getString(cursor.getColumnIndexOrThrow(DBOpenHelper.TB_CREATEDATE));

                //添加到集合
                studentList.add(student);
            }
        }

        cursor.close();

        return studentList;
    }
    /**删除数据*/
    public int delete(Long id) {
        // 获取操作数据库对象
        DB = dBhelpUtil.getWritableDatabase();

        /**
         * String table,  表名
         * String whereClause, 条件
         * String[] whereArgs 参数
         * 返回影响行数，失败 0
         */
        //全部删除
        if(id == null){
            return DB.delete(DBOpenHelper.TABLE_NAME,null,null);
        }
        // 条件查询
        return DB.delete(DBOpenHelper.TABLE_NAME,"id = ?",new String[]{id+""});
    }

    /**保存位图*/
    public void saveBitmap(Wallet student) {
        /** 获取一个写 操作数据的对象*/
        DB = dBhelpUtil.getWritableDatabase();
        //开启事务
        DB.beginTransaction();


        //时间
        Date date = new Date();
        //格式化
        SimpleDateFormat simpleDateFormat = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");

        //执行sql语句 方式
        String sql = "INSERT INTO tb_student(age,sex,name,clazz,createDate,logoHead) VALUES (?,?,?,?,?,?)";
        /**
         * sql 语句
         * 要插入的数据
         */
        DB.execSQL(sql,new Object[]{student.password,student.txt,student.name,student.collect,simpleDateFormat.format(date),student.logoHead});

        //设置事务成功
        DB.setTransactionSuccessful();
        //添加事务
        DB.endTransaction();


    }

    //查询位图
    @RequiresApi(api = Build.VERSION_CODES.P)
    public Wallet selectBitmapById(Long id) {
        /** 获取一个读 操作数据的对象*/
        DB =dBhelpUtil.getReadableDatabase();
        Cursor cursor = null;


        /** 根据id 查询 返回一个游标对象
         * String sql,
         * String[] selectionArgs,
         * select * from tb_student where id = ?
         */
        cursor = DB.rawQuery("select * from "+ DBOpenHelper.TABLE_NAME+" where id =?",new String[]{id+""});
        // 解决报错；android.database.sqlite.SQLiteBlobTooBigException: Row too big to fit into CursorWindow requiredPos=0, totalRows=1
        CursorWindow cw = new CursorWindow("test", 5000000); // 设置CursorWindow的大小为5000000
        AbstractWindowedCursor ac = (AbstractWindowedCursor) cursor;
        ac.setWindow(cw);

        Wallet wallet = null;
        if(cursor != null){
            if(cursor.moveToNext()){
                wallet = new Wallet();
                // 根据游标找到列  在获取数据
                wallet.id = cursor.getLong(cursor.getColumnIndexOrThrow("id"));
                wallet.name = cursor.getString(cursor.getColumnIndexOrThrow(DBOpenHelper.TB_NAME));
                wallet.password = cursor.getString(cursor.getColumnIndexOrThrow(DBOpenHelper.TB_SEX));
                wallet.txt = cursor.getString(cursor.getColumnIndexOrThrow(DBOpenHelper.TB_AGE));
                wallet.collect = cursor.getString(cursor.getColumnIndexOrThrow(DBOpenHelper.TB_CLAZZ));
                wallet.creatDate = cursor.getString(cursor.getColumnIndexOrThrow(DBOpenHelper.TB_CREATEDATE));
                //图片
                wallet.logoHead =cursor.getBlob(cursor.getColumnIndexOrThrow("logoHead")) ;
            }
        }
        cursor.close();

        return wallet;
    }


    //按条件修改
    public int updateById(Wallet student, Long id){
        // 获取写操作数据库对象
        DB = dBhelpUtil.getWritableDatabase();
        //开启事务
        DB.beginTransaction();
        /**
         * String table,
         * ContentValues values, 数据行数据
         * String whereClause, 条件
         * String[] whereArgs   参数
         * 返回影响行数
         */
        //数据行数据
        ContentValues contentValues = new ContentValues();
        contentValues.put(DBOpenHelper.TB_NAME,student.name);
        contentValues.put(DBOpenHelper.TB_SEX,student.password);
        contentValues.put(DBOpenHelper.TB_AGE,student.txt);
        contentValues.put(DBOpenHelper.TB_CLAZZ,student.collect);

        //时间
        Date date = new Date();
        //格式化
        SimpleDateFormat simpleDateFormat = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");
        contentValues.put(DBOpenHelper.TB_CREATEDATE, simpleDateFormat.format(date));

        int result = DB.update(DBOpenHelper.TABLE_NAME,contentValues,"id = ?", new String[]{id+""});
        //完成事务
        DB.setTransactionSuccessful();
        //结束事务
        DB.endTransaction();

        return result;
    }


}