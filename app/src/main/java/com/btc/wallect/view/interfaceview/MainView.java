package com.btc.wallect.view.interfaceview;

import com.btc.wallect.model.entity.Book;
import com.btc.wallect.view.activity.base.BaseConstract;


public interface MainView extends BaseConstract.IBaseView {

    /**
     * 登录成功处理UI
     */
    void setBooksUISuccess(Book book);
}
