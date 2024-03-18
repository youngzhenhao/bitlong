package com.btc.wallect.model.Imoder;

import com.btc.wallect.model.entity.AddMnemonBean;

import java.util.List;

public interface onClickitemToEditListener {
    void onToListData(List<AddMnemonBean> addMnemonlist,int pos);
    void onFocus(int pos);
}
