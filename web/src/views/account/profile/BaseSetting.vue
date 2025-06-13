<template>
  <CollapseContainer :title="t('account.basicSetting')" :canExpand="false">
    <Row :gutter="24">
      <Col :span="14">
        <BasicForm @register="register" />
      </Col>
      <Col :span="10">
        <div class="change-avatar">
          <div class="mb-2">{{ t('account.avatar') }}</div>
          <CropperAvatar
            :uploadApi="uploadApi"
            :value="avatar"
            :btnText="t('account.changeAvatar')"
            :btnProps="{ preIcon: 'ant-design:cloud-upload-outlined' }"
            @change="updateAvatar"
            width="150"
          />
        </div>
      </Col>
    </Row>
    <a-button type="primary" :loading="submitLoading" @click="handleSubmit">
      {{ t('account.updateBasicInfo') }}
    </a-button>
  </CollapseContainer>
</template>
<script lang="ts" setup>
  import { CollapseContainer } from '@/components/Container';
  import { CropperAvatar } from '@/components/Cropper';
  import { BasicForm, useForm } from '@/components/Form';
  import { Col, Row } from 'ant-design-vue';
  import { computed, onMounted, ref } from 'vue';
  import { useI18n } from '@/hooks/web/useI18n';
  // import { useMessage } from '@/hooks/web/useMessage';
  import { getUserInfo, updateBasicInfo } from '@/api/basic/user';
  import { uploadApi } from '@/api/basic/upload';
  import headerImg from '@/assets/images/header.jpg';
  import { useUserStore } from '@/store/modules/user';
  import { baseSettingSchemas } from './data';

  const { t } = useI18n();
  // const { createMessage } = useMessage();
  const userStore = useUserStore();
  let submitLoading: Ref<UnwrapRef<boolean>> = ref(false);
  const [register, { setFieldsValue, validate }] = useForm({
    labelWidth: 120,
    schemas: baseSettingSchemas,
    showActionButtonGroup: false,
  });

  onMounted(async () => {
    const data = await getUserInfo();
    setFieldsValue(data);
  });

  const avatar = computed(() => {
    const { avatar } = userStore.getUserInfo;
    console.log(avatar);
    return avatar || headerImg;
  });

  function updateAvatar({ src, data }) {
    const userinfo = userStore.getUserInfo;
    userinfo.avatar = src;
    userStore.setUserInfo(userinfo);
    console.log('data', data);
  }

  async function handleSubmit() {
    const values = await validate();
    console.log(values);
    submitLoading.value = true;
    updateBasicInfo(values).then(() => {
      submitLoading.value = false;
    });
  }
</script>

<style lang="less" scoped>
  .change-avatar {
    img {
      display: block;
      margin-bottom: 15px;
      border-radius: 50%;
    }
  }
</style>
