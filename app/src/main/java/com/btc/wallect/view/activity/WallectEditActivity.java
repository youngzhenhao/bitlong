package com.btc.wallect.view.activity;

import android.os.Bundle;
import android.view.View;

import com.btc.wallect.R;
import com.btc.wallect.view.activity.base.BaseActivity;

public class WallectEditActivity extends BaseActivity {

    @Override
    protected int setContentView() {
        return R.layout.act_wallect_edit;
    }

    @Override
    protected void init(View view, Bundle savedInstanceState) {
       // setTitle("钱包详情");
        setTitle(getString(R.string.app_txt28));
        setImgBack(true);


    }
}
