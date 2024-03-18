package com.btc.wallect.adapter;

import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.recyclerview.widget.RecyclerView;

import com.btc.wallect.R;
import com.btc.wallect.model.entity.MainTabListBean;

import java.util.List;

public class MainBtcTabAdapter extends RecyclerView.Adapter<MainBtcTabAdapter.ViewHolder> {
    //尾部默认按钮
    private final int DEFAULT_ADD_VIEW = 1;
    //常规布局
    private final int ALBUM_DATA_VIEW = 2;
    private List<MainTabListBean> mFruitList;//用于存放数据的实体类数组


    static class ViewHolder extends RecyclerView.ViewHolder {
        TextView mTvtxttab;
        TextView mTvDownline;

        public ViewHolder(@NonNull View itemView) {
            super(itemView);
            mTvtxttab = itemView.findViewById(R.id.tv_txt_tab);
            mTvDownline = itemView.findViewById((R.id.tv_down_line));

        }

    }

    public MainBtcTabAdapter(List<MainTabListBean> fruitList) {
        mFruitList = fruitList;
    }


    @NonNull
    @Override
    public MainBtcTabAdapter.ViewHolder onCreateViewHolder(@NonNull ViewGroup viewGroup, int i) {
        //动态加载布局
        View view = LayoutInflater.from(viewGroup.getContext()).inflate(R.layout.item_main_btc_tab, viewGroup, false);
        //创建ViewHolder实例，参数为刚加载进来的子项布局
        MainBtcTabAdapter.ViewHolder viewHolder = new MainBtcTabAdapter.ViewHolder(view);//这样，子项布局里面的控件就缓存到了ViewHolder

        return viewHolder;
    }

    @Override
    public void onBindViewHolder(@NonNull ViewHolder holder, int position) {
        MainTabListBean fruit = mFruitList.get(position);      //获取实体类数组中数据
        //将数据bind到子项中控件（子项控件已缓存到了ViewHolder了）
        holder.mTvtxttab.setText(fruit.getTabTxt());


    }


    //作用：返回RecyclerView子项总数
    @Override
    public int getItemCount() {
        return mFruitList.size();
    }
}
