package com.btc.wallect.https;

import com.btc.wallect.model.entity.Book;

import io.reactivex.Observable;
import retrofit2.http.GET;
import retrofit2.http.Query;
/**
 * Created by aiyang on 2018/7/4.
 */

public interface RetrofitAPI {

    @GET("book/search")
    Observable<Book> getSearchBooks(@Query("q") String name,
                                    @Query("tag") String tag, @Query("start") int start,
                                    @Query("count") int count);
}