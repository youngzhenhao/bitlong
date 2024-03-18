package com.btc.wallect.adapter;

import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.recyclerview.widget.RecyclerView;


import com.btc.wallect.R;
import com.btc.wallect.model.entity.CollectBean;

import java.util.List;

public class CollectAdapter extends RecyclerView.Adapter<CollectAdapter.ViewHolder> {
    private List<CollectBean> mFruitList;//用于存放数据的实体类数组


    static class ViewHolder extends RecyclerView.ViewHolder {
        TextView fruitImage;       //用于子项中图片控件的缓存
        TextView fruitName;         //用于子项中名字控件的缓存


        public ViewHolder(@NonNull View itemView) {
            super(itemView);
            fruitImage = itemView.findViewById(R.id.tv_num);
            fruitName = itemView.findViewById((R.id.tv_txt));
        }

    }

    public CollectAdapter(List<CollectBean> fruitList) {
        mFruitList = fruitList;
    }


    @NonNull
    @Override
    public ViewHolder onCreateViewHolder(@NonNull ViewGroup viewGroup, int i) {
        //动态加载布局
        View view = LayoutInflater.from(viewGroup.getContext()).inflate(R.layout.item_collect, viewGroup, false);
        //创建ViewHolder实例，参数为刚加载进来的子项布局
        ViewHolder viewHolder = new ViewHolder(view);//这样，子项布局里面的控件就缓存到了ViewHolder

        return viewHolder;
    }


    public void onBindViewHolder(ViewHolder viewHolder, int i) {
        CollectBean fruit = mFruitList.get(i);      //获取实体类数组中数据
        //将数据bind到子项中控件（子项控件已缓存到了ViewHolder了）
        viewHolder.fruitName.setText(fruit.getName());
        viewHolder.fruitImage.setText(fruit.getImageId()+"");
    }

    //作用：返回RecyclerView子项总数
    @Override
    public int getItemCount() {
        return mFruitList.size();
    }
}

