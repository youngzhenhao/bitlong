package com.btc.wallect.view.activity;

import android.os.Bundle;

import android.view.View;
import android.widget.GridLayout;

import androidx.recyclerview.widget.RecyclerView;
import androidx.recyclerview.widget.StaggeredGridLayoutManager;

import com.btc.wallect.adapter.AddMnemonWordAdapter;
import com.btc.wallect.adapter.ImportSelAdapter;
import com.btc.wallect.model.Imoder.onItemClickListener;
import com.btc.wallect.utils.ConStantUtil;
import com.btc.wallect.utils.DialogUtil;
import com.btc.wallect.view.activity.base.BaseActivity;
import com.btc.wallect.R;
import com.btc.wallect.model.entity.AddMnemonBean;
import com.btc.wallect.model.entity.CollectBean;
import com.btc.wallect.utils.GsonUtils;
import com.btc.wallect.utils.SharedPreferencesHelperUtil;

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


    private  List<AddMnemonBean> inputList1 ;
    private List<CollectBean> collectList2 ;

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
        SharedPreferencesHelperUtil.getInstance().init(this);
        initCollect();//初始化数据
        initAddCollect();
        setRecyclerView1();
        setRecyclerView2();


    }

    private void setRecyclerView1() {
        StaggeredGridLayoutManager mLayoutManager = new StaggeredGridLayoutManager(4, StaggeredGridLayoutManager.VERTICAL);
        mLayoutManager.setOrientation(GridLayout.VERTICAL);
        recyclerView1.setLayoutManager(mLayoutManager);
         fruitAdapter = new AddMnemonWordAdapter(inputList1);
        recyclerView1.setAdapter(fruitAdapter);
    }

    private void setRecyclerView2() {
        StaggeredGridLayoutManager mLayoutManager = new StaggeredGridLayoutManager(5, StaggeredGridLayoutManager.VERTICAL);
        mLayoutManager.setOrientation(GridLayout.VERTICAL);
        recyclerView2.setLayoutManager(mLayoutManager);

        ImportSelAdapter fruitAdapter = new ImportSelAdapter(collectList2);
        recyclerView2.setAdapter(fruitAdapter);
        fruitAdapter.setonItemClickListener(new onItemClickListener() {
            @Override
            public void onItemClick(int position, String txt) {
                String addJson=SharedPreferencesHelperUtil.getInstance().getStringValue(ConStantUtil.ADD_MNEMON,"");
                inputList1=GsonUtils.jsonToList(addJson,AddMnemonBean.class);
                for (int i = 0; i < inputList1.size(); i++) {
                    if (!inputList1.get(i).isState()) {
                       inputList1.get(i).setCollect(txt);
                       inputList1.get(i).setState(true);
                        break;
                    }
                }
                setRecyclerView1();
                SharedPreferencesHelperUtil.getInstance().putStringValue(ConStantUtil.ADD_MNEMON, GsonUtils.listTojson(inputList1));
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
            collectBean.setIndex(i+1);
            collectBean.setState(false);
            inputList1.add(collectBean);
            SharedPreferencesHelperUtil.getInstance().putStringValue(ConStantUtil.ADD_MNEMON, GsonUtils.listTojson(inputList1));
        }
    }

    private void initAddCollect() {
        for (int i = 0; i < 24; i++) {
            CollectBean collectBean = new CollectBean("Collect", i + 1);
            collectList2.add(collectBean);
        }
        Collections.shuffle(collectList2, new Random(System.currentTimeMillis()));
    }


    @OnClick({R.id.tv_sure})
    public void onClick(View view) {
        if (view.getId() == R.id.tv_sure) {
         //   DialogUtil.showSimpleDialog(this, "提示", "确认", null);
            openActivity(MainActivity.class);
        }
    }

}
