package com.btc.wallect.view.activity;

import android.os.Bundle;

import android.view.View;
import android.widget.GridLayout;

import androidx.recyclerview.widget.RecyclerView;
import androidx.recyclerview.widget.StaggeredGridLayoutManager;

import com.btc.wallect.adapter.CollectAdapter;
import com.btc.wallect.utils.HorizontalDividerItemDecoration;
import com.btc.wallect.view.activity.base.BaseActivity;
import com.btc.wallect.R;
import com.btc.wallect.model.entity.CollectBean;
import com.btc.wallect.utils.DialogUtil;

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
        for (int i = 0; i < 24; i++) {
            CollectBean collectBean = new CollectBean("collect", i + 1);
            fruitList.add(collectBean);
        }
    }

    @OnClick({R.id.tv_hand_copy, R.id.tv_cloud_copy, R.id.tv_later_copy})
    public void onClick(View view) {
        if (view.getId() == R.id.tv_hand_copy) {
            openActivity(ImportMnemonicWordActivity.class);
        } else if (view.getId() == R.id.tv_cloud_copy) {
           // DialogUtil.showSimpleDialog(this, "提示", "云备份", null);
            openActivity(MainActivity.class);
        } else if (view.getId() == R.id.tv_later_copy) {
          //  DialogUtil.showSimpleDialog(this, "提示", "稍后备份", null);
            openActivity(MainActivity.class);
        }
    }


}
