package com.btc.wallect.view.activity.fragment;

import android.os.Bundle;

import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.TextView;


import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;

import com.btc.wallect.R;

import butterknife.BindView;
import butterknife.ButterKnife;

/**
 * 消息内容页
 */
public class MsgContentFragment extends Fragment {
    @BindView(R.id.txt_content)
    TextView tvContent;

    private String name;

    @Override
    public void onCreate(@Nullable Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);

        Bundle bundle = getArguments();
        name = bundle.getString("name");
        if (name == null) {
            name = "参数非法";
        }
    }

    @Override
    public View onCreateView(LayoutInflater inflater, ViewGroup container,
                             Bundle savedInstanceState) {
        View view = inflater.inflate(R.layout.fragment_msg_content, container, false);
        ButterKnife.bind(this, view);

        tvContent.setText(name);
        return view;
    }

}
