FROM chmikata/gh-report-cli:0.1.0
COPY entrypoint.sh /app
ENTRYPOINT ["/app/entrypoint.sh"]
CMD ["--help"]
