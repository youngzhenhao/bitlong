package com.btc.wallect.view.activity;

import android.os.Bundle;

import android.util.Log;
import android.view.View;
import android.widget.GridLayout;

import androidx.recyclerview.widget.RecyclerView;
import androidx.recyclerview.widget.StaggeredGridLayoutManager;

import com.btc.wallect.adapter.AddMnemonWordAdapter;
import com.btc.wallect.adapter.ImportSelAdapter;
import com.btc.wallect.model.Imoder.onItemClickListener;
import com.btc.wallect.model.entity.Wallet;
import com.btc.wallect.utils.ConStantUtil;
import com.btc.wallect.utils.DialogUtil;
import com.btc.wallect.utils.LogUntil;
import com.btc.wallect.utils.ToastUtils;
import com.btc.wallect.view.activity.base.BaseActivity;
import com.btc.wallect.R;
import com.btc.wallect.model.entity.AddMnemonBean;
import com.btc.wallect.model.entity.CollectBean;
import com.btc.wallect.utils.GsonUtils;
import com.btc.wallect.utils.SharedPreferencesHelperUtil;
import com.google.gson.Gson;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.Random;

import butterknife.BindView;
import butterknife.OnClick;

public class ImportMnemonicWordActivity extends BaseActivity {
    @BindView(R.id.recycler_view1)
    RecyclerView recyclerView1;
    @BindView(R.id.recycler_view2)
    RecyclerView recyclerView2;


    private List<AddMnemonBean> inputList1;//Fill data
    private List<AddMnemonBean> collectList2;//Fill source data
    private List<CollectBean> sqlCollectList;//Currently stored data

    private AddMnemonWordAdapter fruitAdapter;

    @Override
    protected int setContentView() {
        return R.layout.act_input_mnemon_word;
    }

    @Override
    protected void init(View view, Bundle savedInstanceState) {
        setTitle(" ");
        setImgBack(true);
        inputList1 = new ArrayList<>();
        collectList2 = new ArrayList<>();
        sqlCollectList = new ArrayList<>();
        SharedPreferencesHelperUtil.getInstance().init(this);
        getWallectList();
        initCollect();//初始化数据

        setRecyclerView1();
        setRecyclerView2();


    }

    private void setRecyclerView1() {
        StaggeredGridLayoutManager mLayoutManager = new StaggeredGridLayoutManager(4, StaggeredGridLayoutManager.VERTICAL);
        mLayoutManager.setOrientation(GridLayout.VERTICAL);
        recyclerView1.setLayoutManager(mLayoutManager);
        fruitAdapter = new AddMnemonWordAdapter(inputList1, sqlCollectList);
        recyclerView1.setAdapter(fruitAdapter);
    }

    private void setRecyclerView2() {
        Collections.shuffle(collectList2, new Random(System.currentTimeMillis()));
        StaggeredGridLayoutManager mLayoutManager = new StaggeredGridLayoutManager(5, StaggeredGridLayoutManager.VERTICAL);
        mLayoutManager.setOrientation(GridLayout.VERTICAL);
        recyclerView2.setLayoutManager(mLayoutManager);

        ImportSelAdapter fruitAdapter = new ImportSelAdapter(collectList2);
        recyclerView2.setAdapter(fruitAdapter);
        fruitAdapter.setonItemClickListener(new onItemClickListener() {
            @Override
            public void onItemClick(int position, String txt) {
                String addJson = SharedPreferencesHelperUtil.getInstance().getStringValue(ConStantUtil.ADD_MNEMON, "");
                inputList1 = GsonUtils.jsonToList(addJson, AddMnemonBean.class);
                setVerify(txt);
//                for (int i = 0; i < inputList1.size(); i++) {
//                    if (!inputList1.get(i).isState()) {
//                        LogUntil.e("y元数据：" + sqlCollectList.get(i).getName() + "新数据：" + txt);
//                        if (sqlCollectList.get(i).getName().equals(txt)) {
//                            inputList1.get(i).setCollect(txt);
//                            inputList1.get(i).setState(true);
//                            break;
//
//                        }
//
//                    }
//                    if (!inputList1.get(i).isState()) {
//                        inputList1.get(i).setCollect(txt);
//                        inputList1.get(i).setState(true);
//                        break;
//                    }
//                }
                setRecyclerView1();

            }

            @Override
            public void onItemClick() {

            }
        });


    }


    private void initCollect() {
        for (int i = 0; i < 24; i++) {
            AddMnemonBean collectBean = new AddMnemonBean();
            collectBean.setCollect("");
            collectBean.setIndex(i + 1);
            collectBean.setState(false);
            inputList1.add(collectBean);
            SharedPreferencesHelperUtil.getInstance().putStringValue(ConStantUtil.ADD_MNEMON, GsonUtils.listTojson(inputList1));
        }
    }


    @OnClick({R.id.tv_sure})
    public void onClick(View view) {
        if (view.getId() == R.id.tv_sure) {
            int postion = SharedPreferencesHelperUtil.getInstance().getIntValue(ConStantUtil.WALLECT_VERIFY, 0);

            if (postion > 24) {
                ToastUtils.showToast(this, R.string.app_txt83);
                return;
            }
            ToastUtils.showToast(this, R.string.app_txt84);
            Wallet wallet = new Wallet();
            wallet.verify = ConStantUtil.TRUE;
            wallectDao.updateVerfy(wallet, SharedPreferencesHelperUtil.getInstance().getLongValue(ConStantUtil.CURRENT_SQL_ID, 0));
            openActivity(MainActivity.class);

        }
    }

    private void getWallectList() {
        Long id = SharedPreferencesHelperUtil.getInstance().getLongValue(ConStantUtil.CURRENT_SQL_ID, 0);

        List<Wallet> walletList = selectDataByID(id);
        sqlCollectList = GsonUtils.jsonToList(walletList.get(0).collect, CollectBean.class);
        for (CollectBean collectBean : sqlCollectList) {
            AddMnemonBean addMnemonBean = new AddMnemonBean();
            addMnemonBean.setCollect(collectBean.getName());
            addMnemonBean.setState(false);
            addMnemonBean.setIndex(collectBean.getImageId());
            collectList2.add(addMnemonBean);
        }
        LogUntil.e(new Gson().toJson(sqlCollectList));

    }

    private void setVerify(String txt) {
        int postion = SharedPreferencesHelperUtil.getInstance().getIntValue(ConStantUtil.WALLECT_VERIFY, 0);
        LogUntil.e("LZ>>>" + postion);
        if (postion < sqlCollectList.size()) {
            String collect = sqlCollectList.get(postion).getName();
            if (collect.equals(txt)) {
                inputList1.get(postion).setCollect(txt);
                inputList1.get(postion).setState(true);
                SharedPreferencesHelperUtil.getInstance().putIntValue(ConStantUtil.WALLECT_VERIFY, postion + 1);
                SharedPreferencesHelperUtil.getInstance().putStringValue(ConStantUtil.ADD_MNEMON, GsonUtils.listTojson(inputList1));
            } else {
                inputList1.get(postion).setCollect(txt);
                SharedPreferencesHelperUtil.getInstance().putStringValue(ConStantUtil.ADD_MNEMON, GsonUtils.listTojson(inputList1));
            }
        }


    }

    @Override
    protected void onDestroy() {
        super.onDestroy();
        SharedPreferencesHelperUtil.getInstance().remove(ConStantUtil.WALLECT_VERIFY);

    }
}
