package com.btc.wallect.view.activity;

import android.os.Bundle;
import android.view.View;
import android.widget.ImageView;

import com.btc.wallect.R;
import com.btc.wallect.db.DBOpenHelper;
import com.btc.wallect.db.DBdao;
import com.btc.wallect.model.entity.Wallet;
import com.btc.wallect.utils.ConStantUtil;
import com.btc.wallect.utils.LogUntil;
import com.btc.wallect.view.activity.base.BaseActivity;
import com.google.gson.Gson;

import java.util.List;

import butterknife.BindView;
import butterknife.OnClick;

public class SelCreateWalletAct extends BaseActivity {
    private DBOpenHelper dBhelpUtil;
    public DBdao wallectDa;

    @BindView(R.id.img_create_wallet)
    ImageView imgCreate;

    @Override
    protected int setContentView() {
        return R.layout.act_create_wallet;
    }

    @Override
    protected void init(View view, Bundle savedInstanceState) {
        setTitle(" ");
        setTitleHide(false);
        getWallectList();

    }

    @OnClick({R.id.img_create_wallet, R.id.img_input_wallet, R.id.img_hardWard_wallet, R.id.img_observer_wallet})
    public void onClick(View view) {
        if (view.getId() == R.id.img_create_wallet) {
            openActivityData(CreateWalletActivity.class, ConStantUtil.V_TOACTION_CREATE,ConStantUtil.STATE_TRUE);
            finish();
        } else if (view.getId() == R.id.img_input_wallet) {
            openActivityData(CreateWalletActivity.class, ConStantUtil.V_TOACTION_INPUT,ConStantUtil.STATE_TRUE);

        } else if (view.getId() == R.id.img_hardWard_wallet) {

        } else if (view.getId() == R.id.img_observer_wallet) {
            openActivity(ImportKeyAcivity.class);
        }

    }


    private void getWallectList() {

        List<Wallet> walletList = selectWallectData();
        LogUntil.d(new Gson().toJson(walletList));
    }

    @Override
    protected void onDestroy() {

        super.onDestroy();
    }
}
