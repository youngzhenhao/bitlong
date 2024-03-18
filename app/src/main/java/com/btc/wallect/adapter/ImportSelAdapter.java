package com.btc.wallect.adapter;

import android.annotation.SuppressLint;

import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.recyclerview.widget.RecyclerView;

import com.btc.wallect.model.Imoder.onItemClickListener;
import com.btc.wallect.R;
import com.btc.wallect.model.entity.CollectBean;

import java.util.List;

public class ImportSelAdapter extends RecyclerView.Adapter<ImportSelAdapter.ViewHolder> {
    private List<CollectBean> mFruitList;
    public onItemClickListener listener;
    public void setonItemClickListener(onItemClickListener listener) {
        this.listener = listener;
    }


    static class ViewHolder extends RecyclerView.ViewHolder {

        TextView fruitName;


        public ViewHolder(@NonNull View itemView) {
            super(itemView);

            fruitName = itemView.findViewById((R.id.tv_txt_));

        }

    }

    public ImportSelAdapter(List<CollectBean> fruitList) {
        mFruitList = fruitList;
    }


    @NonNull
    @Override
    public ImportSelAdapter.ViewHolder onCreateViewHolder(@NonNull ViewGroup viewGroup, int i) {
        //动态加载布局
        View view = LayoutInflater.from(viewGroup.getContext()).inflate(R.layout.item_input_sel, viewGroup, false);
        //创建ViewHolder实例，参数为刚加载进来的子项布局
        ImportSelAdapter.ViewHolder viewHolder = new ImportSelAdapter.ViewHolder(view);//这样，子项布局里面的控件就缓存到了ViewHolder

        return viewHolder;
    }


    public void onBindViewHolder(ImportSelAdapter.ViewHolder viewHolder, int i) {
        CollectBean fruit = mFruitList.get(i);      //获取实体类数组中数据
        //将数据bind到子项中控件（子项控件已缓存到了ViewHolder了）
        viewHolder.fruitName.setText(fruit.getName());
        viewHolder.itemView.setOnClickListener(new View.OnClickListener() {
            @SuppressLint("SuspiciousIndentation")
            @Override
            public void onClick(View view) {
                if(listener != null) {
                    int position = viewHolder.getAdapterPosition();
                    if(position != RecyclerView.NO_POSITION)
                        listener.onItemClick(position,fruit.getName());
                }
            }
        });

    }

    //作用：返回RecyclerView子项总数
    @Override
    public int getItemCount() {
        return mFruitList.size();
    }

}


