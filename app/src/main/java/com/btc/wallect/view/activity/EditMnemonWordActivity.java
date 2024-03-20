package com.btc.wallect.view.activity;

import android.os.Bundle;

import android.view.View;
import android.widget.GridLayout;

import androidx.recyclerview.widget.RecyclerView;
import androidx.recyclerview.widget.StaggeredGridLayoutManager;

import com.btc.wallect.adapter.EditMnemontWordAapter;
import com.btc.wallect.adapter.ImportSelAdapter;
import com.btc.wallect.model.Imoder.onClickitemToEditListener;
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
import java.util.List;

import butterknife.BindView;
import butterknife.OnClick;

public class EditMnemonWordActivity extends BaseActivity {
    @BindView(R.id.recycler_view1)
    RecyclerView recyclerView1;
    @BindView(R.id.recycler_view2)
    RecyclerView recyclerView2;


    private List<AddMnemonBean> inputList1 ;
    private List<AddMnemonBean> collectList2 ;
    public List<AddMnemonBean> mqueryList;
    private EditMnemontWordAapter fruitAdapter;
    private ImportSelAdapter inputSelAdapter;
    private int listIndex=0;
    @Override
    protected int setContentView() {
        return R.layout.act_edit_mnemont_word;
    }

    @Override
    protected void init(View view, Bundle savedInstanceState) {
        setTitle("导入助记词");
        setImgBack(true);
        inputList1 = new ArrayList<>();
        collectList2 = new ArrayList<>();
        mqueryList=new ArrayList<>();
        initCollectTest();//初始化数据
        initAddCollect();
        setRecyclerView1();
        setRecyclerView2();


    }

    private void setRecyclerView1() {
        String addJson=SharedPreferencesHelperUtil.getInstance().getStringValue(ConStantUtil.ADD_MNEMON_EDIT,"");
        List<AddMnemonBean> inputedit= GsonUtils.jsonToList(addJson,AddMnemonBean.class);
        StaggeredGridLayoutManager mLayoutManager = new StaggeredGridLayoutManager(4, StaggeredGridLayoutManager.VERTICAL);
        mLayoutManager.setOrientation(GridLayout.VERTICAL);
        recyclerView1.setLayoutManager(mLayoutManager);
        fruitAdapter = new EditMnemontWordAapter(inputedit,mqueryList);
        recyclerView1.setAdapter(fruitAdapter);
        fruitAdapter.setonItemClickListener(new onClickitemToEditListener() {
            @Override
            public void onToListData(List<AddMnemonBean> addMnemonlist, int pos) {
                collectList2.clear();
                for (int i = 0; i < addMnemonlist.size(); i++) {
                    CollectBean collectBean = new CollectBean(addMnemonlist.get(i).getCollect(), i + 1,false);
                    AddMnemonBean addMnemonBean=new AddMnemonBean();
                    addMnemonBean.setState(addMnemonlist.get(i).isState());
                    addMnemonBean.setCollect(addMnemonlist.get(i).getCollect());
                    addMnemonBean.setIndex(addMnemonlist.get(i).getIndex());
                    collectList2.add(addMnemonBean);

//                    collectList2=addMnemonlist;
                }

                inputSelAdapter.notifyDataSetChanged();
            }

            @Override
            public void onFocus(int pos) {
                listIndex=pos;
            }


        });
    }

    private void setRecyclerView2() {
        StaggeredGridLayoutManager mLayoutManager = new StaggeredGridLayoutManager(5, StaggeredGridLayoutManager.VERTICAL);
        mLayoutManager.setOrientation(GridLayout.VERTICAL);
        recyclerView2.setLayoutManager(mLayoutManager);

        inputSelAdapter = new ImportSelAdapter(collectList2);
        recyclerView2.setAdapter(inputSelAdapter);
        inputSelAdapter.setonItemClickListener(new onItemClickListener() {
            @Override
            public void onItemClick(int position, String txt) {
                String addJson=SharedPreferencesHelperUtil.getInstance().getStringValue(ConStantUtil.ADD_MNEMON_EDIT,"");
                inputList1= GsonUtils.jsonToList(addJson,AddMnemonBean.class);
                inputList1.get(listIndex).setCollect(txt);
                SharedPreferencesHelperUtil.getInstance().putStringValue(ConStantUtil.ADD_MNEMON_EDIT, GsonUtils.listTojson(inputList1));

                setRecyclerView1();
            }

            @Override
            public void onItemClick() {

            }
        });

    }


    private void initCollectTest() {
        for (int i = 0; i < 10; i++) {
            if (i==0){
                AddMnemonBean collectBean = new AddMnemonBean();
                collectBean.setCollect("collect");
                collectBean.setIndex(i+1);
                collectBean.setState(false);
                mqueryList.add(collectBean);
            } else if (i==1) {
                AddMnemonBean collectBean = new AddMnemonBean();
                collectBean.setCollect("ccccc");
                collectBean.setIndex(i+1);
                collectBean.setState(false);
                mqueryList.add(collectBean);
            }else if (i==2) {
                AddMnemonBean collectBean = new AddMnemonBean();
                collectBean.setCollect("ccc22");
                collectBean.setIndex(i+1);
                collectBean.setState(false);
                mqueryList.add(collectBean);
            }else if (i==3) {
                AddMnemonBean collectBean = new AddMnemonBean();
                collectBean.setCollect("c333");
                collectBean.setIndex(i+1);
                collectBean.setState(false);
                mqueryList.add(collectBean);
            }else if (i==4) {
                AddMnemonBean collectBean = new AddMnemonBean();
                collectBean.setCollect("c444");
                collectBean.setIndex(i+1);
                collectBean.setState(false);
                mqueryList.add(collectBean);
            }else if (i==5) {
                AddMnemonBean collectBean = new AddMnemonBean();
                collectBean.setCollect("c55");
                collectBean.setIndex(i+1);
                collectBean.setState(false);
                mqueryList.add(collectBean);
            }



        }
        for (int i = 0; i < 10; i++) {
            AddMnemonBean collectBean = new AddMnemonBean();
            collectBean.setCollect("");
            collectBean.setIndex(i+1);
            collectBean.setState(false);
            inputList1.add(collectBean);
        }
        SharedPreferencesHelperUtil.getInstance().putStringValue(ConStantUtil.ADD_MNEMON_EDIT, GsonUtils.listTojson(inputList1));
    }

    private void initAddCollect() {

//        for (int i = 0; i < 24; i++) {
//            CollectBean collectBean = new CollectBean("Collect", i + 1);
//            collectList2.add(collectBean);
//        }
    }


    @OnClick({R.id.tv_sure})
    public void onClick(View view) {
        if (view.getId() == R.id.tv_sure) {

            openActivity(MainActivity.class);
        }
    }
}
