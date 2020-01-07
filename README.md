# Server pipeline


 Proyecto para agrupar los diferentes .gitlab-ci.yml de los proyectos y hacer mas facil su mantenimiento y no exponer  el gitlab token para el uso del mismo.

## URL  modelo

https://gitlab.com/api/v4/projects/9927708/repository/files/default-ci.yaml?ref=master&private_token=mZg.....

## Env necesarias:

```
export GITLAB_TOKEN=djkfjdkfjdkjd
export GITLAB=https://gitlab.com/api/v4/projects/9927708/repository/files/

```


Para usar yaml del branch master hay que usar el endpoint :

``` /master/nombredelarchivo.yaml ```

Para usar yaml del branch development hay que usar el endpoint:

``` /development/nombredelarchivo.yaml ```
