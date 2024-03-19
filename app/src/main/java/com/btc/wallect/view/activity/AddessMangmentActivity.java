package com.btc.wallect.view.activity;

import android.app.Activity;
import android.os.Bundle;
import android.view.View;
import android.widget.GridLayout;

import androidx.annotation.Nullable;
import androidx.recyclerview.widget.RecyclerView;
import androidx.recyclerview.widget.StaggeredGridLayoutManager;

import com.btc.wallect.R;
import com.btc.wallect.adapter.CollectAdapter;
import com.btc.wallect.model.entity.CollectBean;
import com.btc.wallect.utils.HorizontalDividerItemDecoration;
import com.btc.wallect.view.activity.base.BaseActivity;

import java.util.ArrayList;
import java.util.List;

import butterknife.BindView;
import butterknife.OnClick;

public class AddessMangmentActivity extends BaseActivity {


    private final List<CollectBean> fruitList = new ArrayList<>();

    @Override
    protected int setContentView() {
        return R.layout.act_addess_mangemnt;
    }

    @Override
    protected void init(View view, Bundle savedInstanceState) {


    }


    @OnClick({R.id.img_back})
    public void onClick(View view) {
        if (view.getId() == R.id.img_back) {
            finish();
        }
    }
}
