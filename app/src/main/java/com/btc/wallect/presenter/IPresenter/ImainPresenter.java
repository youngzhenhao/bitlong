package com.btc.wallect.presenter.IPresenter;

import com.btc.wallect.view.activity.base.BaseConstract;
import com.btc.wallect.view.interfaceview.MainView;

/**
 * Created by aiyang on 2018/7/4.
 */

public interface ImainPresenter extends BaseConstract.IBasePersenter<MainView> {

    void getSearchBooks(String name,String tag,int start,int count);

}
