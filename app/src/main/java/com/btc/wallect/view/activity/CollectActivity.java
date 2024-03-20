package com.btc.wallect.view.activity;

import android.os.Bundle;

import android.view.View;
import android.widget.GridLayout;

import androidx.recyclerview.widget.RecyclerView;
import androidx.recyclerview.widget.StaggeredGridLayoutManager;

import com.btc.wallect.adapter.CollectAdapter;
import com.btc.wallect.model.entity.Wallet;
import com.btc.wallect.utils.ConStantUtil;
import com.btc.wallect.utils.CopyUtil;
import com.btc.wallect.utils.HorizontalDividerItemDecoration;
import com.btc.wallect.utils.SharedPreferencesHelperUtil;
import com.btc.wallect.utils.ToastUtils;
import com.btc.wallect.view.activity.base.BaseActivity;
import com.btc.wallect.R;
import com.btc.wallect.model.entity.CollectBean;
import com.btc.wallect.utils.DialogUtil;
import com.google.gson.Gson;

import java.util.ArrayList;
import java.util.List;

import butterknife.BindView;
import butterknife.OnClick;

public class CollectActivity extends BaseActivity {
    @BindView(R.id.recycler_view)
    RecyclerView recyclerView;


    private final List<CollectBean> fruitList = new ArrayList<>();

    @Override
    protected int setContentView() {
        return R.layout.act_collect;
    }

    @Override
    protected void init(View view, Bundle savedInstanceState) {
        setTitle(" ");
        setImgBack(true);


        initFruits();//初始化数据

        StaggeredGridLayoutManager mLayoutManager = new StaggeredGridLayoutManager(4, StaggeredGridLayoutManager.VERTICAL);


        HorizontalDividerItemDecoration decoration = new HorizontalDividerItemDecoration(this);
        mLayoutManager.setOrientation(GridLayout.VERTICAL);
        recyclerView.addItemDecoration(decoration);
        recyclerView.setLayoutManager(mLayoutManager);

        CollectAdapter fruitAdapter = new CollectAdapter(fruitList);//创建适配器对象，并传入数据
        recyclerView.setAdapter(fruitAdapter);//为RecyclerView添加带有数据的适配器，从而传给RecyclerView

    }


    private void initFruits() {
        CollectBean collectBean = null;
        for (int i = 0; i < 24; i++) {

            if (i == 0) {
                collectBean = new CollectBean("coll11", i + 1,false);
            } else if (i == 1) {
                collectBean = new CollectBean("colle22", i + 1,false);
            } else if (i == 1) {
                collectBean = new CollectBean("colle333", i + 1,false);
            } else if (i == 1) {
                collectBean = new CollectBean("colle444", i + 1,false);
            } else {
                collectBean = new CollectBean("colle444", i + 1,false);
            }

            fruitList.add(collectBean);
        }
        Wallet wallet = new Wallet();
        String _collect = new Gson().toJson(fruitList);
        wallet.collect = _collect;
        Long id = SharedPreferencesHelperUtil.getInstance().getLongValue(ConStantUtil.CURRENT_SQL_ID, 1);
        wallectDao.updateCollectById(wallet, id);
    }

    @OnClick({R.id.tv_hand_copy, R.id.tv_cloud_copy, R.id.tv_later_copy})
    public void onClick(View view) {
        if (view.getId() == R.id.tv_hand_copy) {
            copyCollect();
            openActivity(ImportMnemonicWordActivity.class);
        } else if (view.getId() == R.id.tv_cloud_copy) {
            ToastUtils.showToast(this, "开发中...");
            openActivity(MainActivity.class);
        } else if (view.getId() == R.id.tv_later_copy) {
            openActivity(MainActivity.class);
            finish();
        }
    }

    private void copyCollect() {
        StringBuilder stringBuilder = new StringBuilder();
        for (int i = 0; i < fruitList.size(); i++) {
            stringBuilder.append(i + "、" + fruitList.get(i).getName()).append(" ");
        }
        String result = stringBuilder.toString().trim();
        CopyUtil.copyClicks(result);
    }


}
