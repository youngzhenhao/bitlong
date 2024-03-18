package com.btc.wallect.utils;

import android.content.Context;
import android.content.SharedPreferences;
import android.util.Log;

public class SharedPreferencesHelperUtil {
    private final String TAG = this.getClass().getSimpleName();

    // file name
    private static final String PREFERENCES_FILE_NAME = "androidpn_demo";

    // instance
    private static SharedPreferencesHelperUtil INSTANCE = null;

    // shared preference
    private SharedPreferences sp;
    // shared preference editor
    private SharedPreferences.Editor editor;

    // is initialized
    private boolean initialized = false;

    public static final synchronized SharedPreferencesHelperUtil getInstance() {
        if (INSTANCE == null) {
            INSTANCE = new SharedPreferencesHelperUtil();
        }
        return INSTANCE;
    }

    public void init(Context context) {
        this.sp = context.getSharedPreferences(PREFERENCES_FILE_NAME,
                Context.MODE_PRIVATE);
        this.editor = sp.edit();
        initialized = true;
    }

    public boolean isInitialized() {
        return initialized;
    }

    /**
     * Push string value
     *
     * @param key
     * @param value
     */
    public void putStringValue(String key, String value) {

        Log.d(TAG, "--------putStringValue----------");
        Log.d(TAG, "key: " + key);
        Log.d(TAG, "value: " + value);

        editor = sp.edit();
        editor.putString(key, value);
        editor.commit();
    }

    /**
     * Get string value by key
     *
     * @param key
     * @param defaultValue
     * @return
     */
    public String getStringValue(String key, String defaultValue) {
        return sp.getString(key, defaultValue);
    }

    /**
     * Push integer value
     *
     * @param key
     * @param value
     */
    public void putIntValue(String key, int value) {
        editor = sp.edit();
        editor.putInt(key, value);
        editor.commit();
    }

    /**
     * Get integer value by key
     *
     * @param key
     * @param defaultValue
     * @return
     */
    public int getIntValue(String key, int defaultValue) {
        return sp.getInt(key, defaultValue);
    }

    /**
     * Push long value
     *
     * @param key
     * @param value
     */
    public void putLongValue(String key, long value) {
        editor = sp.edit();
        editor.putLong(key, value);
        editor.commit();
    }

    /**
     * Get long value by key
     *
     * @param key
     * @param defaultValue
     * @return
     */
    public long getLongValue(String key, long defaultValue) {
        return sp.getLong(key, defaultValue);
    }

    /**
     * Push boolean value
     *
     * @param key
     * @param value
     */
    public void putBooleanValue(String key, boolean value) {
        editor = sp.edit();
        editor.putBoolean(key, value);
        editor.commit();
    }

    /**
     * @param key
     * @return
     */
    public boolean contains(String key) {
        return sp.contains(key);
    }

    /**
     * @param key
     */
    public void remove(String key) {

        Log.d(TAG, "--------remove----------");

        Log.d(TAG, "key: " + key);

        editor = sp.edit();
        editor.remove(key);
        editor.commit();
    }

    /**
     *
     */
    public void clear() {
        editor = sp.edit();
        editor.clear();
        editor.commit();
    }

    /**
     * Get boolean value by key
     *
     * @param key
     * @param defaultValue
     * @return
     */
    public boolean getBooleanValue(String key, boolean defaultValue) {
        return sp.getBoolean(key, defaultValue);
    }
}

