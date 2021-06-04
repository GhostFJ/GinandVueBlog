// 一些验证规则，如手机号、邮箱号等
const telephoneValidator = (val) => /^1[3|4|5|7|8]\d{9}$/.test(val);

export default {
  telephoneValidator,
};
