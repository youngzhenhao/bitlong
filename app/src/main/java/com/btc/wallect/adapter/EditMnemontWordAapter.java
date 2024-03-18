package com.btc.wallect.adapter;

import android.annotation.SuppressLint;

import android.text.Editable;
import android.text.TextWatcher;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.EditText;
import android.widget.RelativeLayout;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.recyclerview.widget.RecyclerView;

import com.btc.wallect.R;
import com.btc.wallect.model.Imoder.onClickitemToEditListener;
import com.btc.wallect.model.entity.AddMnemonBean;
import com.btc.wallect.utils.FuzzySearchExample;

import java.util.List;

public class EditMnemontWordAapter extends RecyclerView.Adapter<EditMnemontWordAapter.ViewHolder> {
    private List<AddMnemonBean> mFruitList;

    private List<AddMnemonBean> mTestList;
    public onClickitemToEditListener listener;
    public void setonItemClickListener(onClickitemToEditListener listener) {
        this.listener = listener;
    }

    static class ViewHolder extends RecyclerView.ViewHolder {
        TextView tv_index;
        EditText ed_collect;
        RelativeLayout mItemlayout;

        public ViewHolder(@NonNull View itemView) {
            super(itemView);

            tv_index = itemView.findViewById((R.id.tv_txt_index));
            ed_collect = itemView.findViewById((R.id.ed_add_txt));
            mItemlayout = itemView.findViewById(R.id.item_layout);
        }

    }

    public EditMnemontWordAapter(List<AddMnemonBean> fruitList,List<AddMnemonBean> testlist) {
        mFruitList = fruitList;
        mTestList = testlist;
    }


    @NonNull
    @Override
    public EditMnemontWordAapter.ViewHolder onCreateViewHolder(@NonNull ViewGroup viewGroup, int i) {
        //动态加载布局
        View view = LayoutInflater.from(viewGroup.getContext()).inflate(R.layout.item_edit_mnemont_word, viewGroup, false);
        //创建ViewHolder实例，参数为刚加载进来的子项布局
        EditMnemontWordAapter.ViewHolder viewHolder = new EditMnemontWordAapter.ViewHolder(view);//这样，子项布局里面的控件就缓存到了ViewHolder

        return viewHolder;
    }


    public void onBindViewHolder(EditMnemontWordAapter.ViewHolder viewHolder, int i) {

        AddMnemonBean fruit = mFruitList.get(i);      //获取实体类数组中数据
        //将数据bind到子项中控件（子项控件已缓存到了ViewHolder了）
        viewHolder.tv_index.setText(fruit.getIndex() + "");
        viewHolder.ed_collect.setText(fruit.getCollect());
        viewHolder.itemView.setOnClickListener(new View.OnClickListener() {
            @SuppressLint("SuspiciousIndentation")
            @Override
            public void onClick(View view) {

            }
        });
        viewHolder.ed_collect.setOnFocusChangeListener(new View.OnFocusChangeListener() {
            @Override
            public void onFocusChange(View v, boolean hasFocus) {
                if (hasFocus) {
                    viewHolder.mItemlayout.setBackgroundResource(R.drawable.bg_side_5af0);
                    listener.onFocus(i);
                }else {
                    viewHolder.mItemlayout.setBackgroundResource(R.drawable.bg_side_f2f2);
                }
            }
        });
        viewHolder.ed_collect.addTextChangedListener(new TextWatcher() {
            @Override
            public void onTextChanged(CharSequence text, int start, int before, int count) {
                String collectTxt=viewHolder.ed_collect.getText().toString();
                if (!collectTxt.isEmpty()) {
                    List<AddMnemonBean> mResultList=FuzzySearchExample.queryList(mTestList,collectTxt);
                    listener.onToListData(mResultList,i);

                }

            }


            @Override
            public void beforeTextChanged(CharSequence text, int start, int count,int after) {

            }

            @Override
            public void afterTextChanged(Editable edit) {

            }
        });

    }

    //作用：返回RecyclerView子项总数
    @Override
    public int getItemCount() {
        return mFruitList.size();
    }
}
