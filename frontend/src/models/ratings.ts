export interface Rating {
  stars: number;
  comment: string;
}

export interface RatingSummary {
  average_stars: number;
  total_ratings: number;
}

export interface FullRating extends Rating {
  rating_id: string;
}
