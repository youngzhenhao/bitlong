package com.btc.wallect.view.activity;

import android.os.Bundle;
import android.view.View;

import androidx.recyclerview.widget.LinearLayoutManager;
import androidx.recyclerview.widget.RecyclerView;

import com.btc.wallect.adapter.TokenDetailAllAdapter;
import com.btc.wallect.model.entity.WalletListBean;
import com.btc.wallect.utils.dialog.TokenDetailDialog;
import com.btc.wallect.view.activity.base.BaseActivity;
import com.btc.wallect.R;

import java.util.ArrayList;
import java.util.List;

import butterknife.BindView;
import butterknife.OnClick;

public class TokenDetailActivity extends BaseActivity {
    @BindView(R.id.recycler_all_detail)
    RecyclerView recyclerView;


    public List<WalletListBean> walletList;


    @Override
    protected int setContentView() {
        return R.layout.act_token_dateil;
    }

    @Override
    protected void init(View view, Bundle savedInstanceState) {
        setTitle(getString(R.string.app_txt28));
        setImgBack(true);

        walletList = new ArrayList<>();
        initFruits();
        initData();

    }

    private void initData() {
        RecyclerView.LayoutManager layoutManager = new LinearLayoutManager(this);
        recyclerView.setLayoutManager(layoutManager);
        TokenDetailAllAdapter mainBtcAdapter = new TokenDetailAllAdapter(walletList);
        recyclerView.setAdapter(mainBtcAdapter);


    }

    private void initFruits() {
        for (int i = 0; i < 5; i++) {
            if (i == 1) {
                WalletListBean walletListBean = new WalletListBean();
                walletListBean.setWallectAmount(0.222);
                walletListBean.setWallectName("BTC-1" + i);
                walletListBean.setWallectKey("bc362....2dfsvd1");
                walletListBean.setType(001);

                walletList.add(walletListBean);
            } else {
                WalletListBean walletListBean = new WalletListBean();
                walletListBean.setWallectAmount(0.222);
                walletListBean.setWallectName("BTC-1" + i);
                walletListBean.setWallectKey("bc362....2dfsvd1");
                walletListBean.setType(002);

                walletList.add(walletListBean);
            }

        }
    }

    @OnClick({R.id.tv_transfer, R.id.tv_collection, R.id.ll_to_dateil})
    public void onClick(View view) {
        if (view.getId() == R.id.tv_transfer) {
            openActivity(CollectMoreActivity.class);
        } else if (view.getId() == R.id.tv_collection) {
            openActivity(CollectQrCodeActivity.class);
        } else if (view.getId() == R.id.ll_to_dateil) {
            TokenDetailDialog tokenDetailDialog = new TokenDetailDialog(this);
            tokenDetailDialog.show();
        }
    }


}
