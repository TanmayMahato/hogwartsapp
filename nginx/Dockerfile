FROM  ubuntu/nginx
EXPOSE 80

COPY default.conf  /etc/nginx/conf.d/
RUN rm -r /etc/nginx/nginx.conf
COPY nginx.conf /etc/nginx/
CMD ["nginx", "-g", "daemon off;"] 
