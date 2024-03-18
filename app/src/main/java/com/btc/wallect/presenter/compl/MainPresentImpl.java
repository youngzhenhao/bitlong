package com.btc.wallect.presenter.compl;

import android.content.Context;
import android.os.Handler;

import com.btc.wallect.https.APIManager;
import com.btc.wallect.view.activity.base.BaseConstract;
import com.btc.wallect.view.activity.base.BasePresenterimpl;
import com.btc.wallect.view.interfaceview.MainView;
import com.btc.wallect.model.entity.Book;
import com.btc.wallect.presenter.IPresenter.ImainPresenter;

import io.reactivex.Observer;
import io.reactivex.android.schedulers.AndroidSchedulers;
import io.reactivex.disposables.Disposable;
import io.reactivex.schedulers.Schedulers;



public class MainPresentImpl extends BasePresenterimpl<MainView> implements ImainPresenter{


    private Book mBook;

    public MainPresentImpl(Context context , BaseConstract.IBaseView view) {
        mAPIManager = new APIManager(context);
        mView = (MainView) view;

    }

    @Override
    public void getSearchBooks(String name, String tag, int start, int count) {
        mView.showProgress();
        mAPIManager.getSearchBooks(name,tag,start,count)
                .subscribeOn(Schedulers.io())
                .observeOn(AndroidSchedulers.mainThread())
                .compose(mView.<Book>bindToLife())
                .subscribe(new Observer<Book>() {

                               @Override
                               public void onComplete() {
                                   if (mBook != null){
                                       //为了效果，我延迟一下下
                                       new Handler().postDelayed(new Runnable() {
                                           @Override
                                           public void run() {
                                               mView.hideProgress();
                                               mView.setBooksUISuccess(mBook);
                                           }
                                       }, 1000);

                                   }
                               }

                               @Override
                               public void onError(Throwable e) {
                                   e.printStackTrace();
                                   mView.hideProgress();
                                   mView.showFaild("请求失败！！");
                               }

                               @Override
                               public void onSubscribe(Disposable d) {

                               }

                               @Override
                               public void onNext(Book book) {
                                   mBook = book;
                               }
                           }
                );

    }
}
