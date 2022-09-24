# go-node-template

## change tmp_project_name

```
find . -type f -print0 | xargs -0 sed -i -e "s/tmp_project_name/test_sample/g" 
```

## install go dependencies

```
cd backend/src
go mod tidy
```

## install vue3

```
npm install -g @vue/cli
vue create frontend
```

## or install anglar

```
npm install -g @angular/cli
ng new frontend
```

## or install svelte

```
npx degit sveltejs/template frontend
cd frontend
npm install
```