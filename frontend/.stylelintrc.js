module.exports = {
  extends: [
    'stylelint-config-recommended',
    'stylelint-config-recommended-scss',
  ],
  overrides: [
    {
      files: ['*.vue', '**/*.vue'],
      extends: ['stylelint-config-recommended-vue'],
    },
  ],
}
