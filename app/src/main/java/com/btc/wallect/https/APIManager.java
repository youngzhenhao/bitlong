package com.btc.wallect.https;

import android.content.Context;

import com.btc.wallect.model.entity.Book;

import io.reactivex.Observable;



public class APIManager{
    /**
     * RetrofitAPI对象
     */
    private RetrofitAPI mRetrofitService;

    /**
     * 构造APIManager，实例化RetrofitAPI
     * @param context
     */
    public APIManager(Context context){
        this.mRetrofitService = RetrofitHelper.getInstance(context).getRetrofitAPI();
    }


    /**
     * 发起请求书类
     * @param name 书名
     * @param tag  关键字
     * @param start 开始页
     * @param count 总页
     * @return
     */
    public Observable<Book> getSearchBooks(String name, String tag, int start, int count){
        return mRetrofitService.getSearchBooks(name,tag,start,count);
    }
}
