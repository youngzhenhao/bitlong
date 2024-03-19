package com.btc.wallect.db;

import android.content.Context;
import android.database.sqlite.SQLiteDatabase;
import android.database.sqlite.SQLiteOpenHelper;


import androidx.annotation.Nullable;

import com.btc.wallect.view.activity.WallectEditActivity;


public class DBOpenHelper extends SQLiteOpenHelper {
    /**数据库名字*/
    public static final String DB_NAME = "walletDB";

    /**表字段信息*/
    public static final String TABLE_NAME = "tb_wallet";
    public static final String TB_NAME = "name";
    public static final String TB_SEX = "password";
    public static final String TB_AGE = "txt";
    public static final String TB_CLAZZ = "collect";
    public static final String TB_CREATEDATE = "createDate";
    public static final String TB_ISSHOW_WALLECT = "show";
    public static final String TB_BTC_KEY = "btcKey";
    public static final String TB_BTC_AMOUNT = "btcAmount";

    /**数据版本号 第一次运行要打开 */
//    public static final int DB_VERSION = 1;

    //模拟数据版本升级
    public static final int DB_VERSION = 2;



    /**
     *
     * @param context   上下文
     * @param name      数据库名字
     * @param factory   游标工厂 null
     * @param version   自定义的数据库版本
     */
    public DBOpenHelper(@Nullable Context context, @Nullable String name, @Nullable SQLiteDatabase.CursorFactory factory, int version) {
        super(context, name, factory, version);
    }

    //数据库第一次创建时被调用
    @Override
    public void onCreate(SQLiteDatabase db) {
        //初始化 第一次 创建数据库
        StringBuilder sql = new StringBuilder();
        sql.append(" create table tb_wallet(");
        sql.append(" id integer primary key,  ");
        sql.append(" name varchar(20),");
        sql.append(" password varchar(2),");
        sql.append(" txt varchar(20),");
        sql.append(" collect varchar(20),");
        sql.append(" createDate varchar(23),");
        sql.append(" show varchar(23),");
        sql.append(" btcKey varchar(23),");
        sql.append(" btcAmount varchar(23) )");

//        Log.e("TAG","------"+sql.toString());

        //执行sql
        db.execSQL(sql.toString());
    }

    //版本号发生改变时调用
    @Override
    public void onUpgrade(SQLiteDatabase db, int oldVersion, int newVersion) {
        //更新数据库 插入字段
        String sql = "alter table tb_student add logoHead varchar(200)";

        db.execSQL(sql);

    }


}