// 用户登录的异步任务
    public class UserLoginTask extends AsyncTask<Void, Void, Result<User>> {

        private final String mName;
        private final String mPassword;

        UserLoginTask(String email, String password) {
            mName = email;
            mPassword = password;
        }

        @Override
        protected Result<User> doInBackground(Void... params) {

            OkHttpClient httpclient = new OkHttpClient();
            MediaType JSON = MediaType.parse("application/json; charset=utf-8");

            User user = new User();
            user.setUserName(mName);
            user.setPassword(mPassword);
            Gson gson = new Gson();
            //将对象转换为诶JSON格式字符串
            String jsonStr = gson.toJson(user);
            RequestBody body = RequestBody.create(JSON, jsonStr);
            Request request = new Request.Builder()
                    .url(SystemParameter.ip + "/user/login")
                    .post(body)
                    .build();
            try {
                Log.d("Request:", request.toString());
                Response response = httpclient.newCall(request).execute();
                String data = response.body().string();

                Log.d("Response:", data);
                return parseJSonWithGSON(data);
            } catch (IOException e) {
                e.printStackTrace();
            }
            return null;
        }

        private Result<User> parseJSonWithGSON(final String data) {
            Gson gson = new Gson();
            Result<User> result = gson.fromJson(data, new TypeToken<Result<User>>() {
            }.getType());
            return result;
        }

        @Override
        protected void onPostExecute(final Result<User> result) {
            mAuthTask = null;
            dismissDialog();

            if (result != null && result.getData() != null) {
                SharedPreferences.Editor editor = sp.edit();
                //保存用户名和密码
                editor.putString("user_name", mName);
                editor.putString("password", mPassword);
                //是否记住密码
                editor.putBoolean("remember_pass", cbRememberPass.isChecked());
                //是否自动登录
                editor.putBoolean("auto_login", autologin.isChecked());
                if (autologin.isChecked())
                    editor.putBoolean("direct_login", true);
                editor.apply();

                SystemParameter.user = result.getData();
                Intent intent = new Intent(LoginActivity.this, MainActivity.class);
                Bundle bundle = new Bundle();
                bundle.putString("username", mName);
                intent.putExtras(bundle);
                startActivity(intent);
                LoginActivity.this.finish();    // 登录成功之后不能回退
            } else {
                password_edit.setError(result == null ? "登录异常" : result.getMessage());
                password_edit.requestFocus();
            }
        }

        @Override
        protected void onCancelled() {
            mAuthTask = null;
            dismissDialog();
        }
    }
	